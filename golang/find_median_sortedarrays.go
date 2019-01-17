// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getkth(a []int, b []int, k int) int {
	m := len(a)
	n := len(b)
	if m > n {
		return getkth(b, a, k)
	}

	if m == 0 {
		return b[k-1]
	}

	if k == 1 {
		return min(a[0], b[0])
	}

	i := min(m, k/2)
	j := min(n, k/2)
	if a[i-1] > b[j-1] {
		return getkth(a, b[j:], k-j)
	} else {
		return getkth(a[i:], b, k-i)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	l := (m + n + 1) >> 1
	r := (m + n + 2) >> 1

	return float64((getkth(nums1, nums2, l) + getkth(nums1, nums2, r))) / 2.0
}

func main() {
	a := []int{1, 2}
	b := []int{3, 9}
	fmt.Println(findMedianSortedArrays(a, b))
}
