package tree

import "fmt"

/**
面向对象
go语言仅支持封装，不支持继承和多态
go语言没有class只有struct
*/

/**
结构创建在堆上还是栈上？
不需要知道,由编译器+运行环境决定
*/
// 结构
type treeNode struct {
	value       int
	left, right *treeNode
}

/**
值接受者 vs 指针接收者
要改变内容必须使用指针接收者
结构过大也考虑使用指针接收者
一致性: 如有指针接收者最好都是指针接收者

值接收者是go语言特有
值/指针接收者均可接收值/指针
*/

// 为结构定义方法
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 只有指针才可以改变结构内容
func (node *treeNode) setValue(value int) {
	// nil 指针也可以调用方法
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored")
		return
	}
	node.value = value
}

// 中序遍历
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 工厂函数
func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}
