package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

// 切片
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:]=", s1)
	s2 := arr[:]
	fmt.Println("arr[:]=", s2)

	fmt.Println("After updateSlice")
	updateSlice(s1)
	updateSlice(s2)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("------------------")
	/*
		1. s1的值为[2,3,4,5] s2的值为 [5,6]
		2. slice 可以向后扩展, 不可向前扩展
		3. s[i] 不可以超越len[s] 向后扩展不可以超越底层数组cap(s)
	*/
	ss1 := arr[2:6]
	ss2 := s1[3:5]

	fmt.Printf("ss1=%v,len(ss1)=%d,cap(ss1)=%d\n",
		ss1, len(ss1), cap(ss1))
	fmt.Printf("ss2=%v,len(ss2)=%d,cap(ss2)=%d\n",
		ss2, len(ss2), cap(ss2))
	/**
	1. 添加元素时，如果超过cap 系统会重新分配更大的底层数组
	2. 由于值传递的关系，必须接受append的返回值
	3. s = append(arr,value)
	*/
	s3 := append(ss2, 10)
	s4 := append(ss2, 10)
	s5 := append(ss2, 10)
	fmt.Println("s3,s4,s5", s3, s4, s5)
}
