package main

import (
	"fmt"
)

func countRangeSum(nums []int, lower int, upper int) int {
	var mergeSort func(int, int) int
	aux := make([]int, len(nums)+1)
	sum := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	mergeSort = func(lo, hi int) int {
		if lo >= hi {
			return 0
		}

		mid := lo + (hi-lo)/2
		res := mergeSort(lo, mid) + mergeSort(mid+1, hi)

		m, n := mid+1, mid+1
		for i := lo; i <= mid; i++ {
			for m <= hi && sum[m]-sum[i] < lower {
				m++
			}
			for n <= hi && sum[n]-sum[i] <= upper {
				n++
			}
			res += n - m
		}

		// inplace merge
		i, j := lo, mid+1
		for k := lo; k <= hi; k++ {
			aux[k] = sum[k]
		}

		for k := lo; k <= hi; k++ {
			switch {
			case i > mid:
				sum[k] = aux[j]
				j++
			case j > hi:
				sum[k] = aux[i]
				i++
			case aux[j] < aux[i]:
				sum[k] = aux[j]
				j++
			default:
				sum[k] = aux[i]
				i++
			}

		}
		return res
	}

	x := mergeSort(0, len(nums))
	return x
}

func main() {
	a := []int{0, 0}
	fmt.Println(countRangeSum(a, 0, 0))
}
