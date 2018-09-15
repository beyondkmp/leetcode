package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findLength(A []int, B []int) int {
	var result int
	d := make([][]int, len(A))
	for i := range d {
		d[i] = make([]int, len(B))
	}

	for i := 0; i < len(A); i++ {
		if A[i] == B[0] {
			d[i][0] = 1
			result = 1
		}
	}

	for i := 0; i < len(B); i++ {
		if B[i] == A[0] {
			d[0][i] = 1
			result = 1
		}
	}

	for i := 1; i < len(A); i++ {
		for j := 1; j < len(B); j++ {
			if A[i] == B[j] {
				d[i][j] = d[i-1][j-1] + 1
				result = max(result, d[i][j])
			} else {
				d[i][j] = 0
			}
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 2, 1}
	b := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(a, b))

	a = []int{0, 0, 0, 0, 0}
	b = []int{0, 0, 0, 0, 0}
	fmt.Println(findLength(a, b))

	a = []int{0, 1, 1, 1, 1}
	b = []int{1, 0, 1, 0, 1}
	fmt.Println(findLength(a, b))
}
