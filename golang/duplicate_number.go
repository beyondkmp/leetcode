/* Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
 *
 * Note:
 * You must not modify the array (assume the array is read only).
 * You must use only constant, O(1) extra space.
 * Your runtime complexity should be less than O(n2).
 * There is only one duplicate number in the array, but it could be repeated more than once. */

package main

import "fmt"

func findDuplicate(nums []int) int {
	saved := make(map[int]int)
	for _, v := range nums {
		saved[v]++
		if saved[v] > 1 {
			return v
		}
	}
	return 0
}

func main() {
	a := []int{1, 1}
	fmt.Println(findDuplicate(a))
}
