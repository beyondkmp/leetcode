package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var firstNode, secondNode, preNode *TreeNode
	var traverse func(head *TreeNode)

	traverse = func(head *TreeNode) {
		if head == nil {
			return
		}

		traverse(head.Left)

		if preNode != nil && preNode.Val >= head.Val {
			if firstNode == nil {
				firstNode = preNode
			}
			secondNode = head
		}
		preNode = head

		traverse(head.Right)
	}
	traverse(root)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}
