package main

import "fmt"

func thirdMax(nums []int) int {
	var MinUint uint
	MaxUint := ^MinUint
	MaxInt := int(MaxUint >> 1)
	MinInt := ^MaxInt

	first, second, third := MinInt, MinInt, MinInt
	for _, n := range nums {
		if n > first {
			third = second
			second = first
			first = n
		} else if n > second && n < first {
			third = second
			second = n
		} else if n > third && n < second {
			third = n
		}
	}
	if third == MinInt {
		return first
	}
	return third
}

func main() {
	x := []int{1, 2, 2, 2, 2, 3}
	fmt.Println(thirdMax(x))
}
