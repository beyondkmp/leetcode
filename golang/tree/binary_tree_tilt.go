package main

/**
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func sumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumTree(root.Left) + sumTree(root.Right) + root.Val
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := sumTree(root.Left)
	rightSum := sumTree(root.Right)
	return abs(leftSum-rightSum) + findTilt(root.Left) + findTilt(root.Right)
}
