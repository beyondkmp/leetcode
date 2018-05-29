package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	cnt := make([]int, len(nums))
	lengest := make([]int, len(nums))
	var result, maxLen int

	for i := 0; i < len(nums); i++ {
		cnt[i], lengest[i] = 1, 1
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				if lengest[i] == lengest[j]+1 {
					cnt[i] += cnt[j]
				}
				if lengest[i] < lengest[j]+1 {
					cnt[i] = cnt[j]
					lengest[i] = lengest[j] + 1
				}
			}
		}
		if maxLen == lengest[i] {
			result += cnt[i]
		}
		if maxLen < lengest[i] {
			maxLen = lengest[i]
			result = cnt[i]
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 4, 3, 5, 4, 7, 2}
	fmt.Println(findNumberOfLIS(a))
}
