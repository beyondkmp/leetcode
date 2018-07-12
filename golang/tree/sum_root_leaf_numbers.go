package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var result int
	var traverse func(h *TreeNode, sum int)

	traverse = func(h *TreeNode, sum int) {
		if h.Left == nil && h.Right == nil {
			result += sum*10 + h.Val
			return
		}

		if h.Left != nil {
			traverse(h.Left, sum*10+h.Val)
		}
		if h.Right != nil {
			traverse(h.Right, sum*10+h.Val)
		}
	}

	traverse(root, 0)
	return result
}

func main() {
	var head TreeNode
	// head.Val = 1
	// head.Left = &TreeNode{Val: 2}
	// head.Right = &TreeNode{Val: 3}

	// fmt.Println(sumNumbers(&head))

	head.Val = 0
	head.Left = &TreeNode{Val: 1}
	fmt.Println(sumNumbers(&head))
}
