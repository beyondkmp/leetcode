package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(pos int, sum int, sol []int)

	sort.Ints(candidates)

	var s []int
	dfs = func(pos int, sum int, sol []int) {
		if sum == target {
			tmp := make([]int, len(sol))
			copy(tmp, sol)
			result = append(result, tmp)
		}

		for i := pos; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				continue
			}
			sum += candidates[i]
			sol = append(sol, candidates[i])
			dfs(i, sum, sol)
			sum -= candidates[i]
			sol = sol[:len(sol)-1]
		}
	}
	dfs(0, 0, s)
	return result
}

func main() {
	c := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(c, target))
}
