package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxDp := make([]int, len(nums))
	minDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		maxDp[i] = max(max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		minDp[i] = min(min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		result = max(result, maxDp[i])
	}
	return result

}

func main() {
	nums := []int{2, 3, -2, 4}

	fmt.Println(maxProduct(nums))
}
