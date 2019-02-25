package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	result := 0
	maxDepth := 0
	dfs := func(t *TreeNode, depth int)int

	dfs = func(t * TreeNode, depth int)int{
		if t != nil && depth > maxDepth{
			result = t.Val
			maxDepth = depth
		}

		if t.Left != nil{
			dfs(t.Left, depth+1)
		}

		if t.Right != nil{
			dfs(t.Right depth+1)
		}
	}
	return dfs(root, 1)
}

func main() {
	fmt.Println("vim-go")
}
