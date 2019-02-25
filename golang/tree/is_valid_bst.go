package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode) (int, int, bool)
	dfs = func(root *TreeNode) (int, int, bool) {
		if root.Left == nil && root.Right == nil {
			return root.Val, root.Val, true
		}
		if root.Left != nil && root.Right != nil {
			leftMin, leftMax, leftValid := dfs(root.Left)
			rightMin, rightMax, rightValid := dfs(root.Right)
			if leftValid && rightValid && leftMax < root.Val && rightMin > root.Val {
				return leftMin, rightMax, true
			}
			return 0, 0, false
		}

		if root.Left != nil {
			leftMin, leftMax, leftValid := dfs(root.Left)
			if leftValid && leftMax < root.Val {
				return leftMin, root.Val, true
			}
			return 0, 0, false
		}

		if root.Right != nil {
			rightMin, rightMax, rightValid := dfs(root.Right)
			if rightValid && rightMin > root.Val {
				return root.Val, rightMax, true
			}
			return 0, 0, false
		}
		return 0, 0, false

	}
	_, _, ok := dfs(root)
	return ok
}
