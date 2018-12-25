package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSwap(A []int, B []int) int {
	swap := make([]int, len(A))
	fix := make([]int, len(A))

	swap[0] = 1
	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > A[i-1] && A[i] > B[i-1] && B[i] > B[i-1] && B[i] > A[i-1]:
			fix[i] = min(swap[i-1], fix[i-1])
			swap[i] = min(swap[i-1], fix[i-1]) + 1
		case A[i] > A[i-1] && B[i] > B[i-1]:
			fix[i] = fix[i-1]
			swap[i] = swap[i-1] + 1
		default:
			fix[i] = swap[i-1]
			swap[i] = fix[i-1] + 1
		}
	}
	fmt.Println(swap, fix)
	return min(swap[len(A)-1], fix[len(A)-1])
}

func main() {
	a := []int{1, 3, 5, 4}
	b := []int{1, 2, 3, 7}
	fmt.Println(minSwap(a, b))
}
