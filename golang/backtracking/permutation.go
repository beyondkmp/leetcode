package main

import "fmt"

func permutation(nums []int, flag []int, pos int, length int) {
	if pos == length {
		fmt.Println(nums)
	} else {
		for i := 1; i <= length; i++ {
			if flag[i] == 0 {
				flag[i] = 1
				nums[pos] = i
				permutation(nums, flag, pos+1, length)
				flag[i] = 0
			}

		}
	}

}
func main() {
	n := 4
	nums := make([]int, n)
	flag := make([]int, n+1)
	permutation(nums, flag, 0, n)

}
