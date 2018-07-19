package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var p2, p3, p5 int
	d := make([]int, n)
	d[0] = 1
	for i := 1; i < n; i++ {
		d[i] = min(d[p2]*2, min(d[p3]*3, d[p5]*5))
		if d[i] == d[p2]*2 {
			p2++
		}
		if d[i] == d[p3]*3 {
			p3++
		}
		if d[i] == d[p5]*5 {
			p5++
		}
	}
	return d[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
