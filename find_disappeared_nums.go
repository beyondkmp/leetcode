package main

import "fmt"

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// 1. 循环数组，将每个数组值(如果已经被改成负的，就算其绝对值)当作索引，将索引项的值改为对应的负值
// 2. 再循环数组，值大于0的索引值就是没有出现过的数字
func findDisappearedNumbers(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	for i, _ := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	test_case := []int{4, 3, 2, 7, 8, 2, 3, 1}
	test_result := findDisappearedNumbers(test_case)
	fmt.Println(test_result)
}
