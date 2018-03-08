package main

import "fmt"

func sortColors(nums []int) {
	var red, white, blue int
	for _, v := range nums {
		switch v {
		case 0:
			red++
		case 1:
			white++
		}
	}
	white += red
	blue = len(nums)
	for i, _ := range nums {
		switch {
		case i < red:
			nums[i] = 0
		case i < white:
			nums[i] = 1
		case i < blue:
			nums[i] = 2
		}
	}
}
func main() {
	nums := []int{0, 1, 0, 2, 1, 1, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}
