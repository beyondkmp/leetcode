/* Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 * Input:nums = [1,1,1], k = 2
 * Output: 2 */

package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int)
	var sum, count int
	preSum[0] = 1
	for _, n := range nums {
		sum += n
		if v, ok := preSum[sum-k]; ok {
			count += v
		}
		if v, ok := preSum[sum]; ok {
			preSum[sum] = v + 1
		} else {
			preSum[sum] = 1
		}
	}
	return count
}

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	fmt.Println(subarraySum(nums, k))
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	k = 2
	fmt.Println(subarraySum(nums, k))
}
