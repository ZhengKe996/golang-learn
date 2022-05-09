package tree

import "fmt"

/**
值接受者 vs 指针接收者
要改变内容必须使用指针接收者
结构过大也考虑使用指针接收者
一致性: 如有指针接收者最好都是指针接收者

值接收者是go语言特有
值/指针接收者均可接收值/指针
*/

// 为结构定义方法
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
