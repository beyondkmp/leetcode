package main

import "fmt"

/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeStack struct {
	items []*TreeNode
	count int
}

func (s *TreeStack) IsEmpty() bool {
	if s.count == 0 {
		return true
	}
	return false
}

func (s *TreeStack) Push(t *TreeNode) {
	s.count++
	s.items = append(s.items, t)
}

func (s *TreeStack) Pop() *TreeNode {
	if s.count == 0 {
		return nil
	}
	item := s.items[s.count-1]
	s.count--
	s.items = s.items[:s.count]
	return item
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	preStack := &TreeStack{}
	preStack.Push(root)

	for !preStack.IsEmpty() {
		item := preStack.Pop()
		result = append(result, item.Val)
		if item.Right != nil {
			preStack.Push(item.Right)
		}
		if item.Left != nil {
			preStack.Push(item.Left)
		}
	}
	return result
}

func main() {
	var head TreeNode
	head.Val = 1
	head.Left = &TreeNode{Val: 2}
	head.Right = &TreeNode{Val: 3}

	fmt.Println(preorderTraversal(&head))
}
