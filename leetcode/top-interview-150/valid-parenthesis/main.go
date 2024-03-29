package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top *StackNode
}

type StackNode struct {
	value    string
	nextElem *StackNode
}

// isEmpty
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

// pop
func (s *Stack) pop() (string, error) {
	if s.isEmpty() {
		return "", errors.New("stack is empty")
	}
	poppedNode := s.top
	s.top = poppedNode.nextElem
	return poppedNode.value, nil
}

// push
func (s *Stack) push(item string) {
	elem := &StackNode{
		value:    item,
		nextElem: s.top,
	}
	s.top = elem
}

// peek
func (s *Stack) peek() (string, error) {
	if s.isEmpty() {
		return "", errors.New("stack is empty")
	}
	return s.top.value, nil
}

const OpeningBracket = 1
const ClosingBracket = 0
const SimpleBrackets = 2
const SquareBrackets = 3
const CurlyBrackets = 4

func checkBracked(s string) (int, error) {
	switch {
	case s == "(" || s == "[" || s == "{":
		return OpeningBracket, nil
	case s == ")" || s == "]" || s == "}":
		return ClosingBracket, nil
	default:
		return -1, errors.New("not a parenthesis")
	}
}

func checkMatchingBrackets(b string, s *Stack) (bool, error) {
	node, err := s.peek()
	if err != nil {
		return false, err
	}
	switch {
	case b == ")":
		return node == "(", nil
	case b == "]":
		return node == "[", nil
	case b == "}":
		return node == "{", nil
	default:
		return false, errors.New("invalid parenthesis")
	}
}

func isValid(s string) bool {
	stack := Stack{
		top: nil,
	}
	// iterate on all the string
	for _, p := range s {
		// get the direction and check if it's a valid parenthesis
		if direction, err := checkBracked(string(p)); err != nil {
			return false
		} else {
			// if it's an opening bracket, put it on the stack
			if direction == OpeningBracket {
				stack.push(string(p))
				// otherwise check if the last parenthesis on the stack matches the current type
			} else {
				if check, err := checkMatchingBrackets(string(p), &stack); err != nil {
					return false
				} else {
					// if it's the same, all good, just pop it from the stack
					if check {
						stack.pop()
						// else, mismatched parethesis, we can already return false
					} else {
						return false
					}
				}
			}
		}
	}
	// check if we iterated over all the parenthesis in the stack, return false if there are still parenthesis on the stack
	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("{[]}"))
}
