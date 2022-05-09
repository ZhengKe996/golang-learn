package tree

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
type Node struct {
	Value       int
	Left, Right *Node
}
