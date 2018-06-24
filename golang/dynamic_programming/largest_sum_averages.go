package main

import (
	"fmt"
	"math"
)

func largestSumOfAverages(A []int, K int) float64 {
	length := len(A)
	if length == 0 {
		return 0
	}

	sum := make([]int, length)
	sum[0] = A[0]
	for i, v := range A[1:] {
		sum[i+1] = sum[i] + v
	}

	if K <= 1 {
		return float64(sum[length-1]) / float64(length)
	}
	if K >= length {
		return float64(sum[length-1])
	}

	dp := make([][]float64, length)
	for i := range dp {
		dp[i] = make([]float64, K)
		dp[i][0] = float64(sum[i]) / float64(i+1)
	}

	for k := 1; k < K; k++ {
		for i := k; i < length; i++ {
			for j := i - 1; j >= k-1; j-- {
				dp[i][k] = math.Max(dp[i][k], dp[j][k-1]+float64(sum[i]-sum[j])/float64(i-j))
			}
		}
	}
	return dp[length-1][K-1]
}

func main() {
	a := []int{9, 1, 2, 3, 9}
	fmt.Println(largestSumOfAverages(a, 3))
}
