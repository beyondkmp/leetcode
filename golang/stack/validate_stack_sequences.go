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

func validateStackSequences(pushed []int, popped []int) bool {
	var st stack
	var p int
	for _, v := range pushed {
		st.Put(v)
		for !st.Empty() && st.Peek() == popped[p] {
			st.Pop()
			p++
		}
	}

	if st.Empty() {
		return true
	}

	return false
}

func main() {
	pushed := []int{1, 0}
	popped := []int{1, 0}
	fmt.Println(validateStackSequences(pushed, popped))

}
