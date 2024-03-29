package main

import "fmt"

// import (
// 	"errors"
// )

// type Stack struct {
// 	top *StackNode
// }

// type StackNode struct {
// 	nodePointer *TreeNode
// 	nextElem    *StackNode
// }

// // isEmpty
// func (s *Stack) isEmpty() bool {
// 	return s.top == nil
// }

// // pop
// func (s *Stack) pop() (*TreeNode, error) {
// 	if s.isEmpty() {
// 		return nil, errors.New("stack is empty")
// 	}
// 	poppedNode := s.top
// 	s.top = poppedNode.nextElem
// 	return poppedNode.nodePointer, nil
// }

// // push
// func (s *Stack) push(item *TreeNode) {
// 	elem := &StackNode{
// 		nodePointer: item,
// 		nextElem:    s.top,
// 	}
// 	s.top = elem
// }

// // peek
// func (s *Stack) peek() (*TreeNode, error) {
// 	if s.isEmpty() {
// 		return nil, errors.New("stack is empty")
// 	}
// 	return s.top.nodePointer, nil
// }

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	fmt.Println(maxDepth())
}
