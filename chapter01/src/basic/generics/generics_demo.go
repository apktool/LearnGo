package main

import "fmt"

func Index[T comparable](s []T, e T) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func main() {
	s1 := []int{10, 20, 15, -10}
	fmt.Println(Index(s1, 15))

	s2 := []string{"foo", "bar", "baz"}
	fmt.Println(Index(s2, "hello"))
}
