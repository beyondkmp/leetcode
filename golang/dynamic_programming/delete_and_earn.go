package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deleteAndEarn(nums []int) int {
	values := make([]int, 10001)

	for _, v := range nums {
		values[v] += v
	}

	var take, skip int
	for i := 0; i < 9; i++ {
		takei := skip + values[i]
		skipi := max(take, skip)

		take = takei
		skip = skipi
	}
	return max(take, skip)
}

func main() {
	a := []int{2, 2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(a))
}
