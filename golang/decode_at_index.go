package main

import (
	"fmt"
	"strings"
)

func decodeAtIndex(S string, K int) string {
	if K == 1 {
		return S[0:1]
	}
	var word []rune
	var result string

	for _, w := range S {
		if w >= '2' && w <= '9' {
			if word == nil {
				word = []rune(result)
			} else {
				result += (string(word))
			}
			result = strings.Repeat(result, int(w-'0'))
			if len(result) >= K {
				return result[K-1 : K]
			}
			word = nil
		} else {
			word = append(word, w)
		}
	}

	if word != nil {
		result += string(word)
	}
	return result[K-1 : K]
}

func main() {
	a := "a2b3c4d5e6f7g8h9"
	k := 69280
	a = "y959q969u3hb22odq595"
	k = 222280369
	fmt.Println(decodeAtIndex(a, k))
}
