package main

import (
	"fmt"
	"sort"
)

func lowerBound(nums []int, key int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == key {
			return mid
		}
		if nums[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high + 1
}

func lengthOfLIS(nums []int) int {
	var tails []int
	for _, v := range nums {
		j := lowerBound(tails, v)
		if j == len(tails) {
			tails = append(tails, v)
		} else {
			tails[j] = v
		}
	}
	return len(tails)
}

type sortByFirst [][]int

func (sb sortByFirst) Len() int {
	return len(sb)
}
func (sb sortByFirst) Swap(i, j int) {
	sb[i], sb[j] = sb[j], sb[i]
}

func (sb sortByFirst) Less(i, j int) bool {
	if sb[i][0] == sb[j][0] {
		return sb[i][1] > sb[j][1]
	}
	return sb[i][0] < sb[j][0]
}
func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(sortByFirst(envelopes))
	second := make([]int, len(envelopes))
	for i, v := range envelopes {
		second[i] = v[1]
	}
	return lengthOfLIS(second)
}

func main() {
	envelopes := [][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}
	fmt.Println(maxEnvelopes(envelopes))
}
