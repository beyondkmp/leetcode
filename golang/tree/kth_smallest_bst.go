package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int
	var traverse func(head *TreeNode)

	traverse = func(head *TreeNode) {
		if head == nil {
			return
		}
		traverse(head.Left)
		if count += 1; count == k {
			result = head.Val
			return
		}
		traverse(head.Right)
	}

	traverse(root)
	return result
}
