package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100 // TODO: 内部修改外部不变
	for i, v := range arr {
		fmt.Println(i, "___", v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int
	fmt.Println(grid)

	// 数组遍历
	for i, v := range arr3 {
		fmt.Println(i, "___", v)
	}

	printArray(arr1)
	//printArray(arr2) // TODO: 长度不同报错
	printArray(arr3)

	println(arr1, arr3) // 内部修改, 外部不变 因为 func printArray(arr [5]int) 会拷贝数组 不同于其他编程语言

}
