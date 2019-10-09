package main

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func widthOfBinaryTree(root *TreeNode) int {
	result := 1
	var dfs func(head *TreeNode, level int, index int)
	var start []int

	dfs = func(head *TreeNode, level int, index int) {
		if head == nil {
			return
		}

		if level == len(start) {
			start = append(start, index)
		} else {
			result = max(result, index-start[level]+1)
		}

		dfs(head.Left, level+1, index*2)
		dfs(head.Right, level+1, index*2+1)
	}

	dfs(root, 0, 1)
	return result
}
