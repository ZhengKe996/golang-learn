package main

import "fmt"

/**
创建；make（make[string]int）
获取元素：m[key]
key不存在时,获得value类型的初始值
用value,ok:=m[key]来判断是否存在key
*/
/**
使用range 遍历key，或者遍历key，value对
不保证遍历顺序，如需顺序需手动对key排序
使用len获得元素个数
*/
func main() {
	m := map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)

	// 遍历
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println("k:", k, "V:", v)
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	fmt.Println(m)
}
