package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binarySearch(nums []int, key int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r + 1
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{Val: val}
	index := binarySearch(preorder[1:], val)

	root.Left = bstFromPreorder(preorder[1 : index+1])
	root.Right = bstFromPreorder(preorder[index+1:])
	return root
}
