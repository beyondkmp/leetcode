package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int, k int) (bool, int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k == 0 {
		return true, sum / k
	}
	return false, 0
}
func canPartitionKSubsets(nums []int, k int) bool {
	ok, avg := canPartition(nums, k)
	if !ok {
		return false
	}

	sort.Ints(nums)

	var dfs func(pos int) bool
	target := make([]int, k)
	for i, _ := range target {
		target[i] = avg
	}

	dfs = func(pos int) bool {
		if pos == -1 {
			return true
		}

		for i := 0; i < k; i++ {
			if target[i] >= nums[pos] {
				target[i] -= nums[pos]
				if dfs(pos - 1) {
					return true
				}
				target[i] += nums[pos]
			}
		}
		return false
	}

	return dfs(len(nums) - 1)
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	fmt.Println(canPartitionKSubsets(nums, k))
}
