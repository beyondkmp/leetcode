package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	var traverse func(head *TreeNode, level int)

	traverse = func(head *TreeNode, level int) {
		if head == nil {
			return
		}

		if level >= len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], head.Val)

		traverse(head.Left, level+1)
		traverse(head.Right, level+1)
	}
	traverse(root, 0)

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}
