package main

import "fmt"

func less(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func partition(a []int, l, r int) int {
	i := l
	j := r + 1
	v := a[l]

	for {

		for {
			i++
			if !less(a[i], v) || i == r {
				break
			}
		}

		for {
			j--
			if !less(v, a[j]) || j == l {
				break
			}
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
	}
	a[j], a[l] = a[l], a[j]
	return j
}

func quicksort(a []int, l, r int) {
	if r <= l {
		return
	}

	i := partition(a, l, r)
	fmt.Println(i)
	fmt.Println(a)
	quicksort(a, l, i-1)
	quicksort(a, i+1, r)
}

func main() {
	a := []int{1, 4, 5, 2, 3, 9}
	quicksort(a, 0, len(a)-1)
	// fmt.Println(partition(a, 0, len(a)-1))
	fmt.Println(a)
}
