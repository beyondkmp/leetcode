package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var head *ListNode
	var tmpHead = new(ListNode)
	head = tmpHead
	for _, v := range a {
		node := ListNode{Val: v, Next: nil}
		tmpHead.Next = &node
		tmpHead = &node
	}
	result := oddEvenList(head.Next)
	for ; result != nil; result = result.Next {
		fmt.Println(result.Val)
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddLast := head
	evenHead := head.Next
	evenLast := evenHead

	for evenLast != nil && evenLast.Next != nil {
		oddLast.Next = evenLast.Next
		oddLast = oddLast.Next
		evenLast.Next = oddLast.Next
		evenLast = evenLast.Next
	}
	oddLast.Next = evenHead

	return head
}
