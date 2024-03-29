package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNext(head, n *ListNode) *ListNode {
	if n.Next == nil {
		return nil // as we don't have a next
	} else {
		n.Next = n.Next.Next
		return head
	}
}

func removeHead(head *ListNode) *ListNode {
	if head.Next == nil {
		// head was the only element
		return nil
	} else {
		head = head.Next
		return head
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	currentElemIndex := 0
	previousNthElem := head
	currentElem := head
	// iterate over all the list and push the position of the nthElement
	for currentElem != nil {
		// intuition: we want to keep one pointer that is n positions behind the currentElem
		// once we start reaching the nth position we start also incrementing the pointer together with the currentElem
		if currentElemIndex > n {
			previousNthElem = previousNthElem.Next
		}
		currentElem = currentElem.Next
		currentElemIndex++
	}
	if currentElemIndex < n {
		// if we scanned less elements than n we canot remove any element
		return nil
	} else if currentElemIndex == n {
		// we should remove the head
		return removeHead(head)
	} else {
		return removeNext(head, previousNthElem)
	}
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	list = removeNthFromEnd(list, 2)
	fmt.Println(list)
}

// 1 2 3 4 5     2
//
