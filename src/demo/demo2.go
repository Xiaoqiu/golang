package main

import (
	"errors"
	"fmt"
)

//error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3 , nil
}
// 自定义的error
type argError struct {
	arg int
	prob string
}
/**
自定义的错误，就是实现系统的内置的这个接口
type error interface {
	Error() string
}
 */
func (e argError) Error() string { //实现接口
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int,error) {
	if arg == 42 {
		return -1, argError{arg,"can't work with it"}
	}
	return arg + 3, nil
}


// Goroutines
func f(from string){
	for i := 0; i < 50; i++ {
		fmt.Println(from,i)
	}
}
func main(){
	for _,i := range []int{7,42} {
		if r,e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		}else{
			fmt.Println("f1 worked: ", r)
		}
	}

	for _,i := range []int{7,42} {
		if r,e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		}else {
			fmt.Println("f2 worked: ", r)
		}
	}

	fmt.Println("=========")
	_, e := f2(42)
	/**
	把实现类输入error,获得属性
	e.(argError)
	 */
	if ae,ok := e.(argError); ok {
		fmt.Println("ok: ", ok)
		fmt.Println("arg: ", ae.arg)
		fmt.Println("prob: ", ae.prob)
	}

	fmt.Println("-----goroutines------")
	//f("direct")
	// 使用go下面两个方法在不同线程异步执行
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")



}

