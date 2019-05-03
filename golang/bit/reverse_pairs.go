package main

import (
	"fmt"
)

func reversePairs(nums []int) int {
	var mergeSort func([]int, int, int) int

	mergeSort = func(a []int, s, e int) int {
		if s >= e {
			return 0
		}

		mid := s + (e-s)/2
		res := mergeSort(a, s, mid) + mergeSort(a, mid+1, e)

		j := mid + 1
		merge := make([]int, e-s+1)
		var k int

		for i, p := s, mid+1; i <= mid; i++ {
			for p <= e && a[i] <= a[p]*2 {
				p++
			}
			res += e - p + 1

			for j <= e && a[i] < a[j] {
				merge[k] = a[j]
				j++
				k++
			}
			merge[k] = a[i]
			k++
		}

		for ; j <= e; j, k = j+1, k+1 {
			merge[k] = a[j]
		}

		copy(nums[s:e+1], merge)
		return res
	}

	return mergeSort(nums, 0, len(nums)-1)
}

func main() {
	a := []int{3, 2, 5, 4, 1}
	fmt.Println(reversePairs(a))
}
