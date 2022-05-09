package main

import (
	"fmt"
	"unicode/utf8"
)

/**
rune 相当于 go的char
使用range遍历pos，rune对
使用utf8.RuneCountInString获得字符数量
使用len获得字节长度
使用[]byte获得字节
*/
func main() {
	s := "啦啦啦啦啦啦啦啦啦啦"
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c) ", i, ch)
	}
	fmt.Println()
}
