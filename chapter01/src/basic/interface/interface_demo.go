package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 将 接口 作为参数传递
func describe1(a Abser) {
	fmt.Printf("(%v, %T)\n", a, a)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
		break
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
		break
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		break
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var v1 Vertex = Vertex{X: 2, Y: 1}
	fmt.Printf("v.X = %.2f; v.Y = %.2f\n", v1.X, v1.Y)

	var v2 Abser = &v1
	fmt.Println(v2.Abs())

	describe1(v2)

	var v3 Abser
	describe1(v3)

	var i1 interface{}
	describe2(i1)

	i1 = 42
	describe2(i1)

	i1 = "Hello World"
	describe2(i1)

	fmt.Println("---------------")

	var i2 interface{} = "Hello World"
	s := i2.(string)
	fmt.Println(s)

	s, ok := i2.(string)
	fmt.Println(s, ok)

	fmt.Println("---------------")

	do(21)
	do("Hello World")
	do(true)

	fmt.Println("---------------")

	a := Person{"Bob", 20}
	b := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, b)
}
