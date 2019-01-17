package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOfthree(a, b, c int) int {
	return min(a, min(b, c))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if word1[m-1] == word2[n-1] {
		return minDistance(word1[:m-1], word2[:n-1])
	}

	return 1 + minOfthree(
		minDistance(word1[:m], word2[:n-1]),
		minDistance(word1[:m-1], word2[:n]),
		minDistance(word1[:m-1], word2[:n-1]),
	)

}

func main() {
	a := "intention"
	b := "execution"
	fmt.Println(minDistance(a, b))

}
