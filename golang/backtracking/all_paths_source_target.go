package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var dfs func(cur int, path []int)

	p := []int{0}

	dfs = func(cur int, path []int) {
		if cur == len(graph)-1 {
			tmp := make([]int, len(path))
			for j := 0; j < len(path); j++ {
				tmp[j] = path[j]
			}
			res = append(res, tmp)
			return
		}

		for _, x := range graph[cur] {
			dfs(x, append(path, x))
		}
	}
	dfs(0, p)
	return res
}

func main() {
	a := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(a))

	b := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(b))
}
