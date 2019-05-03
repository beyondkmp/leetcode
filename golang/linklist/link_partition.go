package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Print(head.Val)
		fmt.Print(" ")
	}
	fmt.Println()
}

func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{}
	more := &ListNode{}
	l, m := less, more

	for ; head != nil; head = head.Next {
		if head.Val < x {
			l.Next, l = head, head
		} else {
			m.Next, m = head, head
		}
	}

	m.Next = nil
	l.Next = more.Next
	return less.Next
}

func main() {
	a := []int{1, 4, 3, 2, 5, 2}
	var head *ListNode
	var tmpHead = new(ListNode)
	head = tmpHead
	for _, v := range a {
		node := ListNode{Val: v, Next: nil}
		tmpHead.Next = &node
		tmpHead = &node
	}
	printLinkList(partition(head.Next, 3))
}
