package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	bfs := []*TreeNode{root}
	i := 0
	for bfs[i] != nil {
		bfs = append(bfs, bfs[i].Left)
		bfs = append(bfs, bfs[i].Right)
		i++
	}
	for _, v := range bfs[i:] {
		if v != nil {
			return false
		}
	}
	return true
}
