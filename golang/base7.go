//Given an integer, return its base 7 string representation.
// Example 1:
// Input: 100
// Output: "202"

// Example 2:
// Input: -7
// Output: "-10"

package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {

	//没有do while, 只能先判断
	if num == 0 {
		return "0"
	}
	var result string
	tmp := num
	if num < 0 {
		tmp = -num
	}
	// result = 求余 + result
	for tmp > 0 {
		result = strconv.Itoa(tmp%7) + result
		tmp /= 7
	}
	if num < 0 {
		result = "-" + result
	}
	return result
}

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(0))
}
