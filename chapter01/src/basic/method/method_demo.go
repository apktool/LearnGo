package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/* function */
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* method */
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 值接收
func (v Vertex) Scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 指针接收
func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc1(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	fmt.Println(Abs(v1))

	fmt.Println("----------------")

	v1.Scale1(2)
	fmt.Printf("X=%.2f, Y=%.2f\n", v1.X, v1.Y)

	v1.Scale2(2)
	fmt.Printf("X=%.2f, Y=%.2f\n", v1.X, v1.Y)

	v2 := Vertex{3, 4}
	ScaleFunc1(v2, 10)
	fmt.Printf("X=%.2f, Y=%.2f\n", v2.X, v2.Y)

	ScaleFunc2(&v2, 10)
	fmt.Printf("X=%.2f, Y=%.2f\n", v2.X, v2.Y)

	fmt.Println("----------------")

	p := &Vertex{3, 4}
	p.Scale1(2)
	fmt.Printf("X=%.2f, Y=%.2f\n", p.X, p.Y)
	ScaleFunc1(*p, 10)
	fmt.Printf("X=%.2f, Y=%.2f\n", p.X, p.Y)
	p.Scale2(2)
	fmt.Printf("X=%.2f, Y=%.2f\n", p.X, p.Y)
	ScaleFunc2(p, 10)
	fmt.Printf("X=%.2f, Y=%.2f\n", p.X, p.Y)

}
