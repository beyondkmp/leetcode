package main

import (
	"fmt"
)

func decodeAtIndex(S string, K int) string {
	var size int

	for _, w := range S {
		if w >= '2' && w <= '9' {
			size *= int(w - '0')
		} else {
			size++
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		w := S[i]

		if K %= size; K == 0 && w >= 'a' && w <= 'z' {
			return S[i : i+1]
		}

		if w >= '2' && w <= '9' {
			size /= int(w - '0')
		} else {
			size--
		}
	}
	return S[0:1]
}

func main() {
	a := "a2b3c4d5e6f7g8h9"
	k := 69280
	a = "y959q969u3hb22odq595"
	k = 222280369
	fmt.Println(decodeAtIndex(a, k))
}
