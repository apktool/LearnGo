package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func for01() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func for02() {
	i := 0
	for ; i < 3; i++ {
		fmt.Println(i)
	}
}

// while
func for03() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

// while True
func for04() {
	var i int = 0
	for {
		if i == 10 {
			fmt.Println("loop end")
			break
		}
		i++
	}
}

// for array
func for05() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	for i, _ := range pow {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func for06() {
	var counter = map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range counter {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}

func if01() {
	var x = math.Sqrt(4)
	if x < 4 {
		fmt.Println("x is less than 4")
	} else {
		fmt.Println("x is equal to 4")
	}
}

func if02() {
	if v := math.Pow(4, 2); v < 20 {
		fmt.Println("v is less than 20")
	}
}

func switch01() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		break
	case "linux":
		fmt.Println("Linux.")
		break
	case "windows":
		fmt.Println("Windows.")
		break
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch02() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
		break
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
		break
	case t.Hour() < 19:
		fmt.Println("Good evening!")
		break
	default:
		fmt.Println("Good!")
	}
}

// defer stack
func defer01() {
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("defer 1")
}

func main() {
	for01()
	for02()
	for03()
	for04()
	for05()
	for06()
	fmt.Println("------end for--------")
	if01()
	if02()
	fmt.Println("------end if---------")
	switch01()
	switch02()
	fmt.Println("------end switch---------")
	defer01()
	fmt.Println("------end defer---------")

}
