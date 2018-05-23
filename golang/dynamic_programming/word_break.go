package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	d := make([]bool, len(s)+1)
	d[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if i >= len(w) && d[i-len(w)] && s[i-len(w):i] == w {
				d[i] = true
			}
		}
	}
	return d[len(s)]
}

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
}
