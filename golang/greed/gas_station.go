package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	var remains, total, start int

	for i, g := range gas {
		remains += g - cost[i]
		if remains < 0 {
			start = i + 1
			total += remains
			remains = 0

		}

	}
	if total+remains < 0 {
		return -1
	}
	return start
}
func main() {
	var gas []int = []int{2, 3, 1}
	var cost []int = []int{3, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
