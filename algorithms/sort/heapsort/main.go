package main

import "fmt"

func sink(nums []int, k, N int) {
	for 2*k+1 <= N {
		i := 2*k + 1
		if i < N && nums[i] < nums[i+1] {
			i++
		}

		if nums[k] >= nums[i] {
			break
		}

		nums[k], nums[i] = nums[i], nums[k]
		k = i
	}
}

func heapSort(nums []int) {
	N := len(nums) - 1
	for i := N / 2; i >= 0; i-- {
		sink(nums, i, N)
	}

	for N > 0 {
		nums[N], nums[0] = nums[0], nums[N]
		sink(nums, 0, N-1)
		N--
	}
}

func findKthLargest(nums []int, k int) int {
	N := len(nums) - 1
	for i := N / 2; i >= 0; i-- {
		sink(nums, i, N)
	}

	for N >= len(nums)-k {
		nums[N], nums[0] = nums[0], nums[N]
		sink(nums, 0, N-1)
		N--
	}
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	heapSort(nums)
	fmt.Println(nums)

	fmt.Println(findKthLargest(nums, 3))
}
