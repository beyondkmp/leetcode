package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	keyboardRow := map[rune]int{
		'q': 0, 'w': 0, 'e': 0, 'r': 0, 't': 0, 'y': 0, 'u': 0, 'i': 0, 'o': 0, 'p': 0, 'a': 1, 's': 1, 'd': 1, 'f': 1, 'g': 1, 'h': 1, 'j': 1, 'k': 1, 'l': 1, 'z': 2, 'x': 2, 'c': 2, 'v': 2, 'b': 2, 'n': 2, 'm': 2,
	}
	var result []string
	for _, w := range words {
		w := strings.ToLower(w)
		initVal := keyboardRow[rune(w[0])]
		for i, a := range strings.ToLower(w) {
			if initVal != keyboardRow[a] {
				break
			}
			if i == len(w)-1 {
				result = append(result, w)
			}
		}
	}
	return result
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))

}
