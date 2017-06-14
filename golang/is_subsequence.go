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

//
func isSubsequence1(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == len(s) {
		return true
	}
	return false
}

func main() {
	s := "acb"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
