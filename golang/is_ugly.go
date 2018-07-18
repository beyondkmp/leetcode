package main

import "fmt"

func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	result := num
	for result != 1 {
		switch {
		case result%2 == 0:
			result = result / 2
		case result%3 == 0:
			result = result / 3
		case result%5 == 0:
			result = result / 5
		default:
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isUgly(6))
}
