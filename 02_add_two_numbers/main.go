package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := ListNode{0, nil}
	carry := 0
	current := &resultNode

	// Handle len(l1) != len(l2)
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry

		// Handle Next Node
		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10

		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// Handle The lastest number.
	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}

	return resultNode.Next
}
