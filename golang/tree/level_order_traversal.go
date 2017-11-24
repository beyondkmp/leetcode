/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	level := []*TreeNode{root}
	for root != nil && len(level) >= 1 {
		var tmp []int
		for _, l := range level {
			tmp = append(tmp, l.Val)
		}
		ans = append(ans, tmp)
		var tmpLevel []*TreeNode
		for _, l := range level {
			if l.Left != nil {
				tmpLevel = append(tmpLevel, l.Left)
			}
			if l.Right != nil {
				tmpLevel = append(tmpLevel, l.Right)
			}
		}
		level = tmpLevel
	}
	return ans
}
