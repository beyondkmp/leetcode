package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func shortestToChar(S string, C byte) []int {
	val := make([]int, len(S))
	for i := range val {
		val[i] = len(S)
	}

	pos := -len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			pos = i
		}
		val[i] = min(val[i], i-pos)
	}

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == C {
			pos = i
		}
		val[i] = min(val[i], abs(pos-i))
	}
	return val
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
