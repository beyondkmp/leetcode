package main

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil && fast.Next.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	mid := midList(head)
	t := &TreeNode{Val: mid.Next.Val}

	rightMid := mid.Next.Next
	mid.Next = nil

	t.Left = sortedListToBST(head)
	t.Right = sortedListToBST(rightMid)
	return t

}
