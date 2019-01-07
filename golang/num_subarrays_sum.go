package main

import "fmt"

func numSubarraysWithSum(A []int, S int) int {
	pSum := make([]int, len(A)+1)
	count := make(map[int]int)
	result := 0

	for i, v := range A {
		pSum[i+1] = pSum[i] + v
	}

	for _, v := range pSum {
		result += count[v]
		count[v+S] += 1
	}

	return result
}

func main() {
	a := []int{1, 0, 1, 0, 1}
	//0,1 1 2 2 3
	S := 2
	fmt.Println(numSubarraysWithSum(a, S))
}
