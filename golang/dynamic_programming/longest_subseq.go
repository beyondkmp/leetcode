package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	d := make([][]int, len(word1)+1)
	for i := range d {
		d[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				continue
			}

			if word1[i-1] == word2[j-1] {
				d[i][j] = d[i-1][j-1] + 1
			} else {
				d[i][j] = max(d[i-1][j], d[i][j-1])
			}
		}
	}
	r := d[len(word1)][len(word2)]
	return len(word1) - r + len(word2) - r
}

func main() {
	s1 := "sea"
	s2 := "eat"
	fmt.Println(minDistance(s1, s2))
}
