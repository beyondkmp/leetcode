package main

import "fmt"

func less(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func partition(a []int, l, r int) int {
	i := l + 1
	j := r
	v := a[l]

	for {
		for ; i <= r && less(a[i], v); i++ {
		}

		for ; j >= l && less(v, a[j]); j-- {
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	a[j], a[l] = a[l], a[j]
	return j
}

func quicksort(a []int, l, r int) {
	if r <= l {
		return
	}

	i := partition(a, l, r)
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
