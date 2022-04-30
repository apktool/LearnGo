package main

import (
	"fmt"
)

var a int

var (
	a1 int
	b1 float32
	c1 string
)

func main() {
	fmt.Printf("a: %d, a1: %d, b1:%f, c1:%s\n", a, a1, b1, c1)

	a1 = 10
	b1 = 2.0
	c1 = "hello go"
	fmt.Printf("a: %d, a1: %d, b1:%f, c1:%s\n", a, a1, b1, c1)

	a = 1
	fmt.Println(a)

	var b = 2
	fmt.Printf("%d\n", b)

	var c float64 = 3.14
	fmt.Printf("%f: %T\n", c, c)

	// 自动推导，只能用在函数内部
	var d int64 = 1
	var e int64 = 2
	fmt.Printf("d[%T]: %d, e[%T]: %d\n", d, d, e, e)

	d, e = e, d
	fmt.Printf("d[%T]: %d, e[%T]: %d\n", d, d, e, e)

	f := d + e
	fmt.Printf("%d=%d+%d\n", f, d, e)

	// const
	const pi float64 = 3.1415926
	fmt.Printf("%f\n", pi)

	const pe = 2.7182818
	fmt.Printf("%f\n", pe)

	const (
		g = 1 // g = 1
		h     // h = 1
		i = 2 // i = 2
		j     // j = 2
	)
	fmt.Println(g, h, i, j)

	// const iota
	// 常量生成器
	const (
		sunday    = iota // sunday = 0
		monday           // monday = 1
		tuesday          // tuesday = 2
		wednesday        // wednesday = 3
		thursday         // thursday = 4
		friday           // friday = 5
		saturday         // saturday = 6
	)

	fmt.Printf("sunday=%d\nmonday=%d\ntuesday=%d\nwednesday=%d\nthursday=%d\nfriday=%d\nsaturday=%d\n",
		sunday, monday, tuesday, wednesday, thursday, friday, saturday,
	)
}
