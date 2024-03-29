package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top *StackNode
}

type StackNode struct {
	nodePointer *TreeNode
	nextElem    *StackNode
}

// isEmpty
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

// pop
func (s *Stack) pop() (*TreeNode, error) {
	if s.isEmpty() {
		return nil, errors.New("stack is empty")
	}
	poppedNode := s.top
	s.top = poppedNode.nextElem
	return poppedNode.nodePointer, nil
}

// push
func (s *Stack) push(item *TreeNode) {
	elem := &StackNode{
		nodePointer: item,
		nextElem:    s.top,
	}
	s.top = elem
}

// peek
func (s *Stack) peek() (*TreeNode, error) {
	if s.isEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.top.nodePointer, nil
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stackP := Stack{}
	stackQ := Stack{}
	// add the root elements
	stackP.push(p)
	stackQ.push(q)
	for !stackP.isEmpty() {
		// remove the first element of the queues
		elemP, _ := stackP.pop()
		elemQ, err := stackQ.pop()
		if err != nil {
			// the tree Q is already finished while P still has nodes, they cannot be the same tree
			return false
		}

		if elemP != nil {
			stackP.push(elemP.Left)
			stackP.push(elemP.Right)

			if elemQ == nil {
				return false
			} else {
				if elemP.Val != elemQ.Val {
					return false
				}
			}
		}
		if elemQ != nil {
			stackQ.push(elemQ.Left)
			stackQ.push(elemQ.Right)
		}

	}
	if !stackQ.isEmpty() {
		return false
	} else {
		return true
	}
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		return isSameTree(root.Left, root.Right)
	}
	return false
}

func main() {
	treeQ := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	treeP := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(treeP, treeQ))
}
