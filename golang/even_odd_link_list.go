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
	var tmp_head = new(ListNode)
	head = tmp_head
	for _, v := range a {
		node := ListNode{Val: v, Next: nil}
		tmp_head.Next = &node
		tmp_head = &node
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
	var odd_head *ListNode
	var odd_last *ListNode
	var even_head *ListNode
	var even_last *ListNode

	odd_head = head
	odd_last = odd_head

	if head.Next != nil {
		even_head = head.Next
		even_last = head.Next

	} else {
		return head

	}

	var tmp_head *ListNode
	tmp_head = even_head.Next
	i := 3
	for ; tmp_head != nil; tmp_head, i = tmp_head.Next, i+1 {
		if i%2 == 0 {
			even_last.Next = tmp_head
			even_last = tmp_head

		} else {
			odd_last.Next = tmp_head
			odd_last = tmp_head

		}

	}
	odd_last.Next = even_head
	even_last.Next = nil
	return odd_head

}
