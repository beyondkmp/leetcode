package main

import (
	"fmt"
	"strconv"
)

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	mark := make(map[string]bool)
	var winRecursive func([]int, int) bool

	allNums := make([]int, maxChoosableInteger)
	for i := 1; i <= maxChoosableInteger; i++ {
		allNums[i-1] = i
	}

	winRecursive = func(nums []int, total int) bool {
		var hash string
		length := 0
		for _, n := range nums {
			hash += strconv.Itoa(n)
			length++
		}
		if v, ok := mark[hash]; ok {
			return v
		}
		if nums[length-1] >= total {
			return true
		}
		for i := 0; i < length; i++ {
			tmp := make([]int, length)
			copy(tmp, nums)
			tmp = append(tmp[:i], tmp[i+1:]...)
			if !winRecursive(tmp, total-nums[i]) {
				mark[hash] = true
				return true
			}
		}
		mark[hash] = false
		return false
	}

	return winRecursive(allNums, desiredTotal)

}
func main() {
	fmt.Println(canIWin(10, 40))
}
