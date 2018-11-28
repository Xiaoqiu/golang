package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// channel synchronization 同步通道
// 使用channel通知，别的goroutine工作做完了
func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//channel directions
//把msg写入pings这个channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//把ping这个channel的消息拿出来，写入pong这个channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

//worker pools
/**
- <-chan是输出的channel
- chan<-是输入的channel
*/
func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {

	/**
	channels: routines之前可以传递数据的管道
	- 默认的channel只能有对应的接受者（<- chan）
	才能发送数据到这个channel(chan <-)
	*/
	fmt.Println("-----channels------")
	messages := make(chan string)
	go func() {
		fmt.Println("i am routines")
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)

	/**
	buffering channels
	- 可以配置一个接收一定量参数列表的channel
	- 可配置接收一定数量的参数
	*/
	fmt.Println("-----buffering channels------")
	messages1 := make(chan string, 2)
	messages1 <- "buffered"
	messages1 <- "channel"
	fmt.Println(<-messages1)
	fmt.Println(<-messages1)

	/**
	channel synchronization 同步通道
	- 同步不同的goroutines的操作
	- 例如一个goroutine阻塞，等待另外一个goroutine结束
	*/
	done := make(chan bool, 1)
	go worker(done) //启动一个异步的goroutine,传一个channel进去
	<-done          // 主线程阻塞直到收到这个channel里面的信号
	// 去掉<- done， 程序会在worker没开始就退出

	fmt.Println("------channel directions-------")
	//channel directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println("----select-----")
	/**
	- 多个channel操作的等待
	- 结合goroutine和channels功能强大
	*/

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//同事等待两个并发的goroutines到来
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	//另一种实现方式
	//msg1 := ""
	//msg2 := ""
	// for {
	//	 select {
	//	 case msg1 = <-c1:
	//		 fmt.Println("received", msg1)
	//	 case msg2 = <-c2:
	//		 fmt.Println("received",msg2)
	//	 }
	//	 if msg1 == "one" && msg2 == "two" {
	//		 break
	//	 }
	// }

	fmt.Println("-----Timeouts------")
	/**
	- 调用外部资源，或者有执行时间限制的时候，timeout很有用
	- 使用channel和select实现很方便
	- select和timeout的模式,通过channel交流的goroutine实现
	- go的很多重要的功能都是基于channel和select
	*/
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // 模拟2秒
		c3 <- "result 1"
	}()
	select {
	case res := <-c3: //获取channel的值
		fmt.Println(res)
	case <-time.After(1 * time.Second): //获取channel的值超时
		fmt.Println("timeout 1")
	}

	c4 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c4 <- "result 2"
	}()
	select {
	case res := <-c4: //获取channel的值
		fmt.Println(res)
	case <-time.After(3 * time.Second): //获取channel的值超时
		fmt.Println("timeout 2")
	}

	fmt.Println("---Non-Blocking Channel Operations---")
	messages2 := make(chan string)
	signals := make(chan bool)

	/**
	非阻塞接收，使用select的default实现：
	- 如果messages2接收到消息，执行case
	- 如果messages2没有接收到消息，执行default
	*/
	select {
	case msg := <-messages2:
		fmt.Println("received message2", msg)
	default:
		fmt.Println("no messages2 received")
	}

	/**
	非阻塞发送，使用select的default实现：
	msg2是无法发送到messages2，因为这个channel没有buffer,
	也没有接收者，所以default执行
	*/
	msg2 := "hi"
	select {
	case messages2 <- msg2:
		fmt.Println("sent message", msg2)
	default:
		fmt.Println("no message sent")
	}

	/**
	 对多个channel实现，非阻塞的接收：
	- 使用多个case实现
	*/
	select {
	case msg := <-messages2:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("---Closing Channels---")
	/**
	Closing Channels: 没有值可以发送到这个channel
	- 体现出接收者把工作做完
	- 当主线程没有工作输入到jobs这个channel，就将它关闭
	- 其中goroutine是worker线程

	*/
	jobs := make(chan int, 5)
	done1 := make(chan bool)

	/**
	- work goroutine
	- 当jobs这channel已经关闭并且所有值都被接收后，more=false
	- 当more=false, 把信息输入到done1这个channel通知主线程
	*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}

	}()

	/**
	- 主线程发送3个job到jobs这个channel
	*/
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")
	//主线程阻塞直到done1这个channel接收到信息，叫醒主线程
	<-done1

	fmt.Println("---Range over Channels---")
	/**
	- 使用for loop遍历channel里面的值
	- for loop 循环完2次会结束，因为我们前面关闭了channel
	- 说明：关闭一个为不空的channel是可以的，
	- 并且关闭后，值还是可以被接收的
	*/
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	fmt.Println("---Timers---")
	/**
	- 指定在未来的一个时间点执行一次
	-
	*/
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // timer.C 是一个channel，当这个channel有个输入就过期
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop() //手动停止这个timer
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	/**
	- 你想在一个时间间隔内重复做一件事
	*/
	fmt.Println("---Tickers---")
	//创建一个500ms的ticker，
	// - 这个ticker有个channel,ticker.C这个channel每500ms就有个值输入
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		// 遍历这个channel的range
		for t := range ticker.C {
			fmt.Println("ticker at ", t)
		}
	}()

	//1600ms之后手动停止这个ticker
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("ticker stopped")

	fmt.Println("---worker pools---")
	/**
	- 使用goroutine和channel实现worker pool
	*/
	//jobs1的channel用来发送工作
	//result的channel用来收集结果
	jobs1 := make(chan int, 100)
	result := make(chan int, 100)
	//启动3个worker
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs1, result)
	}

	//发送5个工作到jobs1这个channel,发送完后关闭
	for j := 1; j <= 5; j++ {
		jobs1 <- j
	}
	close(jobs1)

	//收集所有工作的结果
	for a := 1; a <= 5; a++ {
		<-result
	}

	fmt.Println("---rate limiting---")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	//limiter channel 每200ms接收一个值
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		//阻塞，通过获取limiter channel的值，再执行req
		//即使每200ms才会执行一次req
		<-limiter
		fmt.Println("requst :", req, time.Now())
	}
	burstyLimiter := make(chan time.Time, 3)

	//Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 异步填充这个channel
	go func() {
		//200ms内,允许添加3个值到burstyLimiter channel
		for t := range time.Tick(200 * time.Millisecond) {
			fmt.Println("t : ", t)
			burstyLimiter <- t
		}
	}()

	//请求过来，放入burstyRequest channel
	burstyRequest := make(chan int, 9)
	for i := 1; i <= 9; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	// 9个请求的前3个请求，可以在200毫秒内处理完
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	/**

	Running our program we see the first batch of requests handled once every ~200 milliseconds as desired.
	For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
	*/

	fmt.Println("atomic counters")
	//被多个goroutines使用的原子化的计数器，使用sync/atomic package
	var ops uint64
	//开启50个goroutines每个线程都不断累加这个变量
	for i := 0; i < 50; i++ {
		go func() {
			//for { //改为每个线程都加1，结果刚刚好50说明，这个变量是并发安全的
			//原子化增加方法，传入变量内存地址
			atomic.AddUint64(&ops, 1)
			fmt.Println("tmp ops: ", atomic.LoadUint64(&ops))
			//等待1ms
			time.Sleep(time.Millisecond)
			//}
		}()
	}
	time.Sleep(time.Second)
	//为了安全地使用这个变量，当这个变量被过个goroutines使用时，
	// 我们从当前值复制一个副本
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)

	fmt.Println("---Mutexes---")
	/**
	在被多个goroutines调用时，
	上面使用了atomic operations去管理简单的计数器状态，
	对于更加复杂的情况，我们使用mutex去安全获取数据
	*/

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()

		for w := 0; w < 10; w++ {
			go func() {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}()
		}
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()

	fmt.Println("---Stateful Goroutines---")
	/**
	上面使用显式锁和互斥（mutexes）来实现同步代码块在不同的goroutines之间
	另外一种方式是使用goroutines和channel内置的同步功能，去实现。
	channel-based方式的原理：共享内存只属于一个goroutine所有
	*/

	fmt.Println("---sorting---")
	fmt.Println("---sorting by functions---")

}
