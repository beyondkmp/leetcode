package main

import "fmt"

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(i rune) {
	(*s) = append((*s), i)
}

func (s *Stack) Pop() rune {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	for _, w := range s {
		count[w-'a'] += 1
	}

	var stack Stack
	visited := make([]bool, 26)

	for _, w := range s {
		count[w-'a'] -= 1

		if visited[w-'a'] {
			continue
		}

		for !stack.IsEmpty() && stack.Peek() > w && count[stack.Peek()-'a'] > 0 {
			visited[stack.Pop()-'a'] = false
		}

		stack.Push(w)
		visited[w-'a'] = true
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
