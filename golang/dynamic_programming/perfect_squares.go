package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	d := make([]int, n+1)
	for i := range d {
		d[i] = 1<<63 - 1
	}

	d[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			d[i] = min(d[i-j*j]+1, d[i])
		}
	}
	return d[n]
}

func main() {
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(15))
}
