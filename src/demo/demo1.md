- 常量
- 方法
- Multiple Return Values
- Variadic Functions 可变参的方法
- Closures 闭包函数
- Recursion递归 
- 指针
- structs
- 接口
- 变量
- 常量 constant
- for loop
- if/else
- switch
- array
- slices
- map
- range

#### 常量 
```go
    package main
    import (
        "fmt"
        "math"
        "time"
    )
    
    const s string = "constants"
    func main()  {
        //...程序入口
    }

```

#### 方法

```go
    func plus(a int,b int) int {
        return a + b
    }
    func pluss(a, b, c int) int {
        return a + b + c
    }
func main()  {
	fmt.Println("------function------")
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res1 := pluss(1, 2, 3)
	fmt.Println("1 + 2 + 3 =",res1)
	}
```
#### Multiple Return Values 返回多个变量
```go

    func vals() (int, int){
        return  3,7
    }
func main()  {
	fmt.Println("----Multiple Return Values----")
	a1, b1 := vals()
	fmt.Println(a1)
	fmt.Println(b1)

	_, c1 := vals()
	fmt.Println(c1)
	}
```
#### Variadic Functions 可变参的方法
````go

    func sum1(nums ...int) {
        fmt.Print(nums, " ")
        total := 0
        for _, num := range nums {
            total += num
        }
        fmt.Println(total)
    }
func main()  {
	fmt.Println("---Variadic Functions---")
	sum1(1, 2)
	sum1(1, 2, 3)

	nums1 := []int{1, 2, 3, 4}
	sum1(nums1...)
	}

````
#### Closures 闭包函数:
   内层函数引用了外层函数的变量，其返回值也是一个函数
```go
    func intSeq() func() int {
        i := 0
        return func() int {
            i++
            return i
        }
    }
func main()  {
	fmt.Println("-----Closures-----")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	}
```
#### Recursion递归 
```go
    func fact (n int) int {
        if n == 0 {
            return 1
        }
        return n * fact(n-1)
    }
func main()  {
	fmt.Println("----Recursion----")
	fmt.Println(fact(7))
	}
```
#### 指针

```go

    func zeroval(ival int) {
        ival = 0
    }
    func zeroptr(iptr *int) { // 入参是*int表示int pointer
        *iptr = 0 //* 这个符号，把值从地址解析出来，重新赋值到这个地址
    }
func main()  {
	fmt.Println("-------method------")
	// 定义结构的方法
	r := rect{width:10, height:5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	//使用指针调用是一样的
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	
	fmt.Println("----Pointers-----")
    	i2 := 1
    	fmt.Println("initial: ", i2)
    
    	zeroval(i2)
    	fmt.Println("zeroval: ", i2)
    
    	fmt.Println("address for i2: ", &i2)
    	fmt.Println("value for i2: ", i2)
    
    	//改变了i2这个变量，因为这个函数获取了变量的内存地址
    	zeroptr(&i2) //&表示获取i2变量的内存地址
    	fmt.Println("zeroptr: ", i2)
    
    	fmt.Println("pointer", &i2)
    	}
```

####  structs
```go
    type person struct {
        name string
        age int
    }
func main()  {
    fmt.Println("----structs-----")

	fmt.Println(person{"bob", 20})
	fmt.Println(person{name:"alice", age:20})
	fmt.Println(person{name:"fred"})
	fmt.Println(&person{name:"ann",age:40})
	s := person{name:"sean",age:50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	}
```
####  接口
```go
    type geometry interface {
        area() float64
        perim() float64
    }
    //struct
    type rect struct {
        width, height int
    }
    type circle struct {
        radius float64
    }
    type rect2 struct {
        width, height float64
    }
    //定义结构的方法
    func (r *rect) area() int {
        return r.height * r.width
    }
    func (r rect) perim() int {
        return 2*r.width + 2*r.height
    }
    func (r rect2) area() float64 {
        return r.height * r.width
    }
    func (r rect2) perim() float64 {
        return 2*r.width + 2*r.height
    }
    
    func (c circle) area() float64 {
        return math.Pi * c.radius * c.radius
    }
    func (c circle) perim() float64{
        return 2 * math.Pi * c.radius
    }
    
    func measure(g geometry) {
        fmt.Println("g: ",g)
        fmt.Println("g.area()", g.area())
        fmt.Println("g.perim()", g.perim())
    }
func main()  {
    fmt.Println("----Interfaces-----")
        r2 := rect2{width:3, height:4}
        c2 := circle{radius:5}
        measure(r2)
        measure(c2)
        }
```
#### 变量
```go
func main()  {
	fmt.Println("hello world!")
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
	}
```
#### 常量 constant

```go
func main()  {
	fmt.Println(s)
	const  n  = 500000000
	const d02 = 3e20 / n
	fmt.Println(d02)
	fmt.Println(int(64))
	fmt.Println(math.Sin(n))
	}
```
#### for loop
```go
func main()  {
	fmt.Println("------for loop-------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
    for j := 7; j<= 9; j++ {
        fmt.Println(j)
    }

    for{
        fmt.Println("loop")
        break
    }

    for n := 0; n <=5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
    }
```
#### if/else
````go
func main()  {
	fmt.Println("----------if/else---------")

	if 7%2 == 0 {
		fmt.Println("7 is even")
	}else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9 ; num < 0 {
		fmt.Println(num, "is negative")
	}else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}else {
		fmt.Println(num, "has multiple digits")
	}
	}
````

#### switch

```go
func main()  {
	fmt.Println("---switch----")

	i01 := 2
	fmt.Println("Write", i01, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
```
#### array
```go
func main()  {
    fmt.Println("-------array-------")
	var a02[5] int
	fmt.Println("emp:", a02)

	a02[4] = 100
	fmt.Println("set: ", a02)
	fmt.Println("get: ", a02[4])

	fmt.Println("len: ", len(a02))
	b02 := [5]int{1,2,3,4,5}
	fmt.Println("dcl: ", b02)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3 ;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	}
```
#### slices
```go
func main()  {
    fmt.Println("-------slices-------")
	s02 := make([]string, 3)
	fmt.Println("emp: ", s02)

	s02[0] = "a"
	s02[1] = "b"
	s02[2] = "c"
	fmt.Println("set: ", s02)
	fmt.Println("get: ", s02[2])

	fmt.Println("len: ", len(s02))

	s02 = append(s02, "d")
	s02 = append(s02,"e", "f")
	fmt.Println("apd: ", s02)

	c02 := make([]string, len(s02))
	copy(c02,s02)
	fmt.Println("cpy: ", c02)

	l := s02[2:5]
	fmt.Println("sl1: " , l)

	l = s02[:5]
	fmt.Println("sl2: ", l)

	l = s02[2:]
	fmt.Println("sl3: ", l)

	t02 := []string{"g" ,"h", "i"}
	fmt.Println("dcl: ", t02)

	twoD02 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD02[i] = make([]int, innerLen)
		for j := 0; j < innerLen ; j++ {
			twoD02[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD02)
}
```
#### map
````go
func main()  {
    fmt.Println("-----------map-------------")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))
	delete(m, "k2")
	fmt.Println("map: ",m)
	_,prs := m["k2"]
	fmt.Println("prs: ",prs)

	n2 := map[string]int{"foo":1,"bar":2}
	fmt.Println("map: ", n2)
}
````
#### range
```go
func main()  {
	fmt.Println("------range------")
	nums := []int{2,3,4}
	sum := 0

	/**
	range 返回index, value两个值，
	我们不需要index,所以使用_忽略了这个值
	 */
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum: ",sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v := range kvs {
		fmt.Println("%s -> $s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
	
}
```


