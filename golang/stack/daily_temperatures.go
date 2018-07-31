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
		if st.Empty() {
			st.Put(i)
			continue
		}

		for !st.Empty() {
			if v > temperatures[st.Peek()] {
				result[st.Peek()] = i - st.Peek()
				st.Pop()
			} else {
				break
			}
		}
		st.Put(i)
	}

	return result
}

func main() {
	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(a))
}
