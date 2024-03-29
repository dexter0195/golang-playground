package main

// **
//   - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sum(elem1, elem2 *ListNode, carryOver int) (int, int) {
	sum := 0
	if elem1 == nil {
		sum = elem2.Val
	} else if elem2 == nil {
		sum = elem1.Val
	} else {
		sum = elem1.Val + elem2.Val
	}
	sum += carryOver

	// compute the carry over
	if sum >= 10 {
		carryOver = 1
	} else {
		carryOver = 0
	}
	return sum % 10, carryOver
}

func appendList(head, currentElem *ListNode, val int) (*ListNode, *ListNode) {
	newElem := &ListNode{
		Val:  val,
		Next: nil,
	}
	if head == nil {
		head = newElem
	} else {
		currentElem.Next = newElem
	}
	currentElem = newElem
	return head, currentElem
}

func sumListElements(elem1, elem2, sumList, sumElem *ListNode, carryOver int) (*ListNode, *ListNode, int) {
	// iterate over the list
	currentSum, carryOver := sum(elem1, elem2, carryOver)
	sumList, sumElem = appendList(sumList, sumElem, currentSum)

	return sumList, sumElem, carryOver
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentElem1 := l1
	currentElem2 := l2
	// should we sum in-place in any of the two list or should we create a new list ? let's create a new list
	var sumList *ListNode
	var currentElemSum *ListNode

	carryOver := 0

	for currentElem1 != nil || currentElem2 != nil {
		sumList, currentElemSum, carryOver = sumListElements(currentElem1, currentElem2, sumList, currentElemSum, carryOver)
		if currentElem1 != nil {
			currentElem1 = currentElem1.Next
		}
		if currentElem2 != nil {
			currentElem2 = currentElem2.Next
		}
	}
	if carryOver != 0 {
		sumList, _ = appendList(sumList, currentElemSum, carryOver)
	}

	return sumList
}

func main() {
	var currElem1 *ListNode
	var list1head *ListNode
	number1 := []int{9, 9, 9, 9, 9, 9, 9}

	for _, n := range number1 {
		list1head, currElem1 = appendList(list1head, currElem1, n)
	}

	var currElem2 *ListNode
	var list2head *ListNode
	number2 := []int{9, 9, 9, 9}

	for _, n := range number2 {
		list2head, currElem2 = appendList(list2head, currElem2, n)
	}

	addTwoNumbers(list1head, list2head)

}
