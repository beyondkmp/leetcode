package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	coursesGraph := make(map[int][]int)

	for _, p := range prerequisites {
		coursesGraph[p[1]] = append(coursesGraph[p[1]], p[0])
	}

	visted := make([]bool, numCourses)
	onPath := make([]bool, numCourses)

	var dfsCycle func(nodeNum int) bool
	dfsCycle = func(nodeNum int) bool {
		if visted[nodeNum] {
			return false
		}
		onPath[nodeNum], visted[nodeNum] = true, true

		for _, adjacency := range coursesGraph[nodeNum] {
			if onPath[adjacency] || dfsCycle(adjacency) {
				return true
			}
		}
		onPath[nodeNum] = false
		return false
	}

	for i := 0; i < numCourses; i++ {
		if !visted[i] && dfsCycle(i) {
			return false
		}
	}
	return true
}

func main() {
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(2, prerequisites))
}
