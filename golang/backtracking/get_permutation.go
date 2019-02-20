package main

func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	flag := make([]bool, n)
	var visitAll func(pos int)
	var count int

	visitAll = func(pos int) {
		if pos == n {
			count++
			return
		}

		for i := 1; i <= n; i++ {
			if count == k {
				return
			}
			if flag[i-1] == false {
				nums[pos] = byte(i + 48)
				flag[i-1] = true
				visitAll(pos + 1)
				flag[i-1] = false
			}
		}
	}

	visitAll(0)
	return string(nums)
}
