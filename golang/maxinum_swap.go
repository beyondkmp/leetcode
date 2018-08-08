package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	maxidx := -1
	maxdigit := -1
	leftidx := -1
	rightidx := -1

	for i := len(s) - 1; i >= 0; i-- {
		if int(s[i]) > maxdigit {
			maxdigit = int(s[i])
			maxidx = i
			continue
		}

		if int(s[i]) < maxdigit {
			leftidx = i
			rightidx = maxidx
		}
	}
	if leftidx == -1 {
		return num
	}
	s = s[:leftidx] + s[rightidx:rightidx+1] + s[leftidx+1:rightidx] + s[leftidx:leftidx+1] + s[rightidx+1:]
	num, _ = strconv.Atoi(s)
	return num
}

func main() {
	fmt.Println(maximumSwap(7298))
}
