package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var stack []*TreeNode
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			cur = top.Right
		}
	}
	return result
}
