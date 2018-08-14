package main

import (
	"fmt"
)

func isDigit(w byte) bool {
	if w >= '1' && w <= '9' {
		return true
	}
	return false
}

func isAlpha(w byte) bool {
	if w >= 'a' && w <= 'z' {
		return true
	}
	return false
}

func decodeAtIndex(S string, K int) string {
	var size int

	for i := range S {
		w := S[i]
		if isDigit(w) {
			size *= int(w - '0')
		} else {
			size++
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		w := S[i]

		if K %= size; K == 0 && isAlpha(w) {
			return S[i : i+1]
		}

		if isDigit(w) {
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
