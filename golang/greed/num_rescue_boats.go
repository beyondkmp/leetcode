package main

import (
	"fmt"
	"sort"
)

type byWeight []int

func (w byWeight) Len() int           { return len(w) }
func (w byWeight) Less(i, j int) bool { return w[i] > w[j] }
func (w byWeight) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

func numRescueBoats(people []int, limit int) int {
	flag := make([]bool, len(people))
	result := 0
	var i, j int

	sort.Sort(byWeight(people))
	for i = 0; i < len(people); i++ {
		if flag[i] {
			continue
		}
		for j = i + 1; j < len(people); j++ {
			if !flag[j] && people[i]+people[j] <= limit {
				flag[j] = true
				result++
				break
			}
		}
		if j >= len(people) {
			flag[i] = true
			result++
		}
	}
	return result
}

func numRescueBoats2(people []int, limit int) int {
	result := 0
	high := len(people) - 1
	low := 0

	sort.Sort(sort.IntSlice(people))
	for high >= low {
		if people[high]+people[low] <= limit {
			low++
		}
		high--
		result++
	}
	return result
}

func main() {
	people := []int{3, 5, 3, 4}
	limit := 5
	fmt.Println(numRescueBoats(people, limit))
}
