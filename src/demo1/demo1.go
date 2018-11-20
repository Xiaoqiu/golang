package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constants"
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

	//constant
	fmt.Println(s)
	const  n  = 500000000
	const d02 = 3e20 / n
	fmt.Println(d02)
	fmt.Println(int(64))
	fmt.Println(math.Sin(n))


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

