package tree

import "fmt"

// 只有指针才可以改变结构内容
func (node *Node) SetValue(value int) {
	// nil 指针也可以调用方法
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored")
		return
	}
	node.Value = value
}
