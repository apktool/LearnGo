package main

import "fmt"

func main() {
	var a [4]int = [4]int{2, 3, 4, 6}
	fmt.Println(a)

	var b []string = []string{"a", "b", "c"}
	fmt.Println(b)

	var c [2]string
	c[0] = "Hello"
	c[1] = "World"
	fmt.Println(c)

	// slice
	var d []int = a[1:3]
	fmt.Printf("a=%v, d=%v\n", a, d)

	d[0] = 0
	fmt.Printf("a=%v, d=%v\n", a, d)

	var e = []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(e)

	// make slice
	f := make([]int, 5)
	fmt.Println(f)

	var g []int
	fmt.Println(g == nil)

	g = append(g, 1)
	g = append(g, 3)
	g = append(g, 5)
	fmt.Println(g)

	// for array
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
