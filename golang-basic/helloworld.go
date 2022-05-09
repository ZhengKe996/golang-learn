package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 变量定义
var (
	aa = 3
	ss = "Hello"
	bb = true
)

func varibaleZeroValue() {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)
}

func varibleInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func varibleTypeDeduction() {
	var a, b, c, s = 3, 4, false, "Hello"
	fmt.Println(a, b, c, s)
}

func varibleShorter() {
	a, b, c, s := 3, 4, false, "Hello"
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	fmt.Printf("%0.3f",
		cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
const (
	CC string = "Hello"
	DD bool   = false
)

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 枚举
func enums() {
	const (
		cpp = iota
		java
		pthon
		golang
	)

	fmt.Println(cpp, java, pthon, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello World")
	varibaleZeroValue()
	varibleInitialValue()
	varibleTypeDeduction()
	varibleShorter()

	fmt.Println(aa, ss, bb)

	fmt.Println("-----------------")
	euler()

	triangle()

	consts()

	enums()
}
