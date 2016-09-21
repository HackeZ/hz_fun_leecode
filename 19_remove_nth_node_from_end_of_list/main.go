package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fastNode := head
	slowNode := &ListNode{Val: head.Val, Next: head}
	var flag bool

	for n > 1 {
		fastNode = fastNode.Next
		n--
	}
	for fastNode.Next != nil {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
		flag = true
	}

	slowNode.Next = slowNode.Next.Next
	if flag {
		return head
	}
	return head.Next
}

func showList(head *ListNode) {
	for head.Next != nil {
		fmt.Printf("%d => ", head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}

func main() {
	first := ListNode{Val: 4, Next: nil}
	second := ListNode{Val: 3, Next: &first}
	third := ListNode{Val: 2, Next: &second}
	four := ListNode{Val: 1, Next: &third}

	head := removeNthFromEnd(&four, 1)
	showList(head)
}
