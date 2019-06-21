package main

import "sort"

type BITree []int

func (b BITree) update(i int) {
	for i <= len(b)-1 {
		b[i] += 1
		i += (i & -i)
	}
}

func (b BITree) getSum(i int) int {
	s := 0
	for i != 0 {
		s += b[i]
		i -= (i & -i)
	}
	return s
}

func countSmaller(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)

	rank := make(map[int]int)
	for i, v := range tmp {
		rank[v] = i + 1
	}

	bitTree := make(BITree, len(nums)+1)
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = bitTree.getSum(rank[nums[i]] - 1)
		bitTree.update(rank[nums[i]])
	}

	return result
}
