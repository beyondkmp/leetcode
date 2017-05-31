/* Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.
 *
 * You need to find the shortest such subarray and output its length.
 *
 * Example 1:
 * Input: [2, 6, 4, 8, 10, 9, 15]
 * Output: 5
 * Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
 * Note:
 * Then length of the input array is in range [1, 10,000].
 * The input array may contain duplicates, so ascending order here means <=. */

package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	var low, high int
	var lowFlag, highFlag bool

	low = 0
	high = len(nums) - 1
	for !lowFlag || !highFlag {
		for i := low + 1; i <= high-1; i++ {
			if nums[i] < nums[low] {
				lowFlag = true
			}
			if nums[i] > nums[high] {
				highFlag = true
			}
		}
		if nums[low] > nums[high] {
			return high - low + 1
		}
		if !lowFlag {
			low++
		}
		if !highFlag {
			high--
		}
		if low >= high {
			return 0
		}
	}
	return high - low + 1
}

func main() {
	a := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(a))
	a = []int{7}
	fmt.Println(findUnsortedSubarray(a))
	a = []int{6, 7}
	fmt.Println(findUnsortedSubarray(a))
	a = []int{7, 6}
	fmt.Println(findUnsortedSubarray(a))
	a = []int{1, 2, 3, 3, 3}
	fmt.Println(findUnsortedSubarray(a))

}
