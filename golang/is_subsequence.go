package main

import "fmt"

func minOfValue(nums []int, value int) int {
	min := -1
	flag := 1
	for _, n := range nums {
		if n > value {
			if flag == 1 {
				min = n
				flag = 0
			} else {
				if n < min && n > value {
					min = n
				}
			}
		}
	}
	return min
}

func isSubsequence(s string, t string) bool {
	position := make(map[rune][]int)
	for i, r := range t {
		position[r] = append(position[r], i+1)
	}
	min := 0
	fmt.Println(position)
	for _, r := range s {
		if v, ok := position[r]; ok {
			tmp := minOfValue(v, min)
			fmt.Println(v)
			fmt.Println(tmp, min)
			if tmp < min {
				return false
			}
			min = tmp
		} else {
			return false
		}
	}
	return true
}
func main() {
	s := "acb"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
