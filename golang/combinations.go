package main

import "fmt"

func combine_r(nums []int, pos int, start int, n int, k int) [][]int {
	var results [][]int
	if pos == k {
		tmp := make([]int, k)
		copy(tmp, nums)
		return [][]int{tmp}
	} else {
		for i := start; i <= n; i++ {
			nums[pos] = i
			results = append(results, combine_r(nums, pos+1, i+1, n, k)...)
		}
	}
	return results
}

func combine(n int, k int) [][]int {
	nums := make([]int, 3)
	return combine_r(nums, 0, 1, n, k)
}
func main() {
	fmt.Println(combine(5, 3))
}
