/*
Best Time to Buy and Sell Stock
[7, 1, 5, 3, 6, 4]
[7, 1, 1, 1, 1, 1]  # 到当前位置时的最小值
[0, 0, 4, 2, 5, 3]

*/
package main

import "fmt"

func maxProfit(prices []int) int {
	beforemin := prices[0]
	max := 0
	for _, v := range prices {
		if v < beforemin {
			beforemin = v
		}
		if v-beforemin > max {
			max = v - beforemin
		}
	}
	return max
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
