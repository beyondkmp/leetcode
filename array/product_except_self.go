/*
   [1, 2, 3, 4]
1. [1, 1, 1*2, 1*2*3]
#从后往前循环
2. [4*3*2, 4*3, 4, 1]

队列1 * 队列2 = 最后的结果

*/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	p := 1
	for j := len(nums) - 1; j >= 0; j-- {
		result[j] *= p
		p *= nums[j]

	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
