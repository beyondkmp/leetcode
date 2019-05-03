package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	val := head.Val
	p := head.Next
	if p.Val != val {
		head.Next = deleteDuplicates(p)
		return head
	}

	for p != nil && p.Val == val {
		p = p.Next
	}

	return deleteDuplicates(p)
}
