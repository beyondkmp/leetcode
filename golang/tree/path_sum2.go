package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	var dfs func(r *TreeNode, sum int, sol []int)

	dfs = func(r *TreeNode, sum int, sol []int) {
		if r == nil {
			return
		}

		if r.Left == nil && r.Right == nil && sum == r.Val {
			tmp := make([]int, len(sol)+1)
			copy(tmp, sol)
			tmp = append(tmp, r.Val)
			result = append(result, tmp)
			return
		}

		sol = append(sol, r.Val)
		dfs(r.Left, sum-r.Val, sol)
		dfs(r.Right, sum-r.Val, sol)
	}

	dfs(root, sum, nil)
	return result
}
