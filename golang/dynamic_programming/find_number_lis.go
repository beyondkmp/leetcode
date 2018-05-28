package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	longest := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			longest[i] = longest[i-1] + 1
		} else {
			longest[i] = longest[i-1]
		}
	}
	var result int
	for i := len(nums) - 1; i >= 0; i-- {
		if longest[i] == longest[i-1] {
			result++
		} else {
			break
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 5, 4, 7}
	fmt.Println(findNumberOfLIS(a))
}
