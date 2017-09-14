package main

import (
	"fmt"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	len1 := len(s1)
	len2 := len(s2)

	remainder := 0
	var result []int
	var i, j int
	for i, j = len1-1, len2-1; i >= 0 && j >= 0; {
		sum := (s1[i] + s2[j] + remainder) % 10
		remainder = (s1[i] + s2[j] + remainder) / 10
		result = append(result, sum)
		i--
		j--
	}

	for ; i >= 0; i-- {
		sum := (s1[i] + remainder) % 10
		remainder = (s1[i] + remainder) / 10
		result = append(result, sum)
	}
	for ; j >= 0; j-- {
		sum := (s2[j] + remainder) % 10
		remainder = (s2[j] + remainder) / 10
		result = append(result, sum)
	}

	if remainder > 0 {
		result = append(result, remainder)
	}

	var head ListNode
	flag := &head
	for i = len(result) - 1; i >= 0; i-- {
		var node ListNode
		node.Val = result[i]
		flag.Next = &node
		flag = &node

	}
	return head.Next

}

func convertLiceToList(s []int) *ListNode {
	var head ListNode
	flag := &head
	for _, v := range s {
		var node ListNode
		node.Val = v
		flag.Next = &node
		flag = &node
	}
	return head.Next
}

func printList(l *ListNode) {
	s := ""
	for l != nil {
		if l.Next != nil {
			s = s + strconv.Itoa(l.Val) + "->"
		} else {
			s = s + strconv.Itoa(l.Val)
		}
		l = l.Next
	}
	fmt.Println(s)
}
func main() {
	s1 := []int{7, 2, 4, 3}
	s2 := []int{5, 6, 4}
	l1 := convertLiceToList(s1)
	printList(l1)
	l2 := convertLiceToList(s2)
	printList(l2)

	printList(addTwoNumbers(l1, l2))
}
