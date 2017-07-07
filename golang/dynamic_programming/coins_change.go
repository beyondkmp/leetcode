package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	var visitAll func(amount int) int
	flag := make(map[int]int)

	visitAll = func(a int) int {
		if a < 0 {
			return -1
		}
		if a == 0 {
			return 0
		}
		if v, ok := flag[a]; ok {
			return v
		}

		min := math.MaxInt32
		for _, v := range coins {
			res := visitAll(a - v)
			if res >= 0 && res < min {
				min = res + 1
			}
		}
		if min == math.MaxInt32 {
			min = -1
		}
		flag[a] = min
		return min
	}

	visitAll(amount)
	return flag[amount]
}

func main() {
	/* coins := []int{186, 419, 83, 408}
	 * amount := 6249 */
	coins := []int{2}
	amount := 3
	nums := coinChange(coins, amount)
	fmt.Println(nums)
}
