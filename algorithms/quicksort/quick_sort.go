package main

import "fmt"

func less(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func partition(a []int, l, r int) int {
	i, j := l, r+1
	v := a[l]

	for {
		for i = i + 1; i <= r && less(a[i], v); i++ {
		}

		for j = j - 1; j >= l && less(v, a[j]); j-- {
		}

		if i > j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}

	a[j], a[l] = a[l], a[j]
	return j
}

func partition1(a []int, l, r int) int {
	m := l - 1
	v := a[r]
	for i := l; i <= r; i++ {
		if a[i] <= v {
			m++
			a[m], a[i] = a[i], a[m]
		}
	}
	return m
}

func quicksort(a []int, l, r int) {
	if r <= l {
		return
	}

	i := partition1(a, l, r)
	quicksort(a, l, i-1)
	quicksort(a, i+1, r)
}

func main() {
	a := []int{1, 7, 2, 4, 5, 9, 8, 2, 6, 3, 9}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)

	a = []int{1, 2, 2, 2, 2, 2}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5, 6}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)

	a = []int{1, 1}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
