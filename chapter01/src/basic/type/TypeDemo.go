package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	/**
	 * 有符号整型: int, int8, int32, int64, rune
	 * 无符号整型: uint, uint8, uint32, uint64, uintptr
	 */
	var a int64 = -1
	fmt.Printf("a: %d\n", a)

	var b uint64 = 1
	fmt.Printf("b: %d\n", b)

	/**
	 * 浮点型: float32, float64
	 */
	var c float32 = 1.0
	fmt.Printf("c: %f\n", c)

	var d float64 = 1 << 24
	fmt.Printf("d: %.2f\n", d)

	/**
	 * 布尔
	 */
	var e bool = true
	fmt.Printf("e: %t, 1=1 %t \n", e, 1 == 1)
	fmt.Println(1 == 1)

	/**
	 * 字符类型
	 * byte类型是uint8的别名，rune类型是int32的别名
	 */
	var f byte = 'A'
	fmt.Printf("f: %c\n", f)

	var g byte = 65
	fmt.Printf("g: %c\n", g)

	var h rune = '\u0041'
	fmt.Printf("h: %c, %t \n", h, unicode.IsLetter(h))

	var i string = "hello\nworld"
	fmt.Printf("i: %s, len=%d\n", i, len(i))

	var j string = `hello\nworld`
	fmt.Printf("j: %s, len=%d\n", j, len(j))

	var k string = "hello " + "world"
	fmt.Printf("k: %s, len=%d\n", k, len(k))
	fmt.Printf("k: %s, o in pos %d or pos %d\n", k, strings.Index(k, "o"), strings.LastIndex(k, "o"))

	/**
	 * 类型转换
	 */
	var l float64 = 1.0
	var m int = int(l)
	fmt.Printf("l=%f, m=%d\n", l, m)
	fmt.Printf("int64 range: min=%d, max=%d\n", math.MinInt64, math.MaxInt64)

	n := "localhost:8080"
	o := []byte(n)
	o[len(o)-1] = '1'
	p := string(o)
	fmt.Printf("n=%s, p=%s\n", n, p)
}
