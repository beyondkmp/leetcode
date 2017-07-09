package main

import (
	"fmt"
	"math"
)

//非递归版本
func coinChange(coins []int, amount int) int {
	values := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, v := range coins {
			res := i - v
			if res >= 0 && values[res] >= 0 && values[res] < min {
				min = values[res] + 1

			}

		}
		if min == math.MaxInt32 {
			min = -1

		}
		values[i] = min

	}
	return values[amount]
}

func coinChangeR(coins []int, amount int) int {
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
