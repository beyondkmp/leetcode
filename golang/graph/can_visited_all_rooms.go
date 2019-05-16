package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	var dfs func(r int, visited map[int]bool)
	var res bool
	dfs = func(r int, visited map[int]bool) {
		if len(visited) == len(rooms) {
			res = true
		}
		for _, v := range rooms[r] {
			if res {
				return
			}
			if _, ok := visited[v]; !ok {
				visited[v] = true
				dfs(v, visited)
				visited[v] = false
			}
		}
	}
	v := map[int]bool{0: true}
	dfs(0, v)
	return res
}

func main() {
	a := [][]int{{1}, {2}, {3}, {}}
	fmt.Println(canVisitAllRooms(a))

	a = [][]int{{1, 3}, {1, 0, 3}, {2}, {0}}
	fmt.Println(canVisitAllRooms(a))
}
