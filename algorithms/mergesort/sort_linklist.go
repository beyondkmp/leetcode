package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := midList(head)
	right := mid.Next
	mid.Next = nil
	return merge(sortList(head), sortList(right))
}

func midList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	last := &ListNode{}
	head := last

	for left != nil && right != nil {
		if left.Val <= right.Val {
			last.Next = left
			left = left.Next
		} else {
			last.Next = right
			right = right.Next
		}
		last = last.Next
	}

	if left != nil {
		last.Next = left
	}

	if right != nil {
		last.Next = right
	}

	return head.Next
}

func main() {
	nums := []int{3, 2, 4, 1, 5, 10, 17, 2, -8, 0}
	dummy := &ListNode{}
	head := dummy
	for _, v := range nums {
		tmp := &ListNode{Val: v}
		head.Next = tmp
		head = head.Next
	}
	sortedList := sortList(dummy.Next)
	for sortedList != nil {
		fmt.Println(sortedList.Val)
		sortedList = sortedList.Next
	}
}
