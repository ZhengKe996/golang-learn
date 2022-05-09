package main

import "tree"

func main() {
	var root tree.treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	root.right.left = new(treeNode)

	root.left.right = createTreeNode(2)
	root.print()
}
