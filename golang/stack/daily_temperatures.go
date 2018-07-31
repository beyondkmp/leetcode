package main

import "fmt"

type stack []int

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Peek() int {
	return s[len(s)-1]
}

func (s *stack) Put(i int) {
	(*s) = append((*s), i)
}

func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var st stack

	for i, v := range temperatures {
		for !st.Empty() {
			peek := st.Peek()
			if v <= temperatures[peek] {
				break
			}

			result[peek] = i - peek
			st.Pop()
		}
		st.Put(i)
	}

	return result
}

func main() {
	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(a))
}
