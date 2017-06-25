package main

import "fmt"

func combine(n int, k int) [][]int {
	nums := make([]int, k)
	var results [][]int
	var visitAll func(pos int, start int)

	visitAll = func(pos int, start int) {
		if pos == k {
			tmp := make([]int, k)
			for j := 0; j < k; j++ {
				tmp[j] = nums[j]
			}
			results = append(results, tmp)
			return
		}
		for i := start; i <= n; i++ {
			nums[pos] = i
			visitAll(pos+1, i+1)
		}
	}

	visitAll(0, 1)
	return results
}
func main() {
	fmt.Println(combine(5, 3))
}
