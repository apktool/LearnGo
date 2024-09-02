package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	/* 函数也是值。可以像其他值一样传递。 */
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	var a float64 = hypot(3, 4)
	fmt.Println(a)

	var b float64 = compute(hypot)
	fmt.Println(b)

	var c float64 = compute(math.Pow)
	fmt.Println(c)

	/* 函数也是值。可以用作函数的参数或返回值 */
	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
