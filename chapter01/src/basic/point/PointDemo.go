package main

import "fmt"

/**
 * 指针(pointer)在Go语言中可以被拆分为两个核心概念：
 * 1. 类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无需拷贝数据，类型指针不能进行偏移的运算
 * 2. 切片，由指向起始元素的原始指针，元素数量和容量组成
 */

func main() {
	var a int = 1
	fmt.Printf("&a=%p\n", &a)

	p1 := &a
	fmt.Printf("p1=%p, *p1=%d, type=%T\n", p1, *p1, p1)

	var p2 *int = &a
	fmt.Printf("p2=%p, *p2=%d, type=%T\n", p2, *p2, p2)

	p3 := new(string)
	*p3 = "hello go"
	fmt.Printf("p3=%p, *p3=%s, type=%T\n", p3, *p3, p3)
}
