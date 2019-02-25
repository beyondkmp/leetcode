package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var result []int
	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(result) {
			result = append(result, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	return result
}
