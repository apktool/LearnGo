package main

import "fmt"

func main() {
	var a = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(a)

	delete(a, "a")
	fmt.Println(a)

	var b map[string]int = make(map[string]int)
	b["a"] = 1
	b["b"] = 2
	b["c"] = 3
	fmt.Println(b)

	for k, v := range b {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
