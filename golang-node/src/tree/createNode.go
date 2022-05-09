package tree

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
