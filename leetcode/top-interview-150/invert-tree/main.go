package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		tmp := root.Left
		root.Left = invertTree(root.Right)
		root.Right = invertTree(tmp)
	}
	return root
}

func main() {
	fmt.Println()
}
