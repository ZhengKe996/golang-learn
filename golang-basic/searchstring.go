package main

import "fmt"

/**
寻找最长不含有重复字符的子串
abcabcbb -> abc
bbbbb -> b
pwwkew -> wke
*/

/**
对于每一个字母 x
	lastOccurred[x]不存在 或者 < start -> 无需操作
	lastOccurred[x] >= start -> 更新 start
	更新 lastOccurred[x] 更新 maxLength
*/
// 不支持中文
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 支持中文
func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))

	fmt.Println(lengthOfNonRepeatingSubStr2("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr2("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr2("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr2("你好啊啊"))
}
