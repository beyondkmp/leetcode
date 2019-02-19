package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitsAlphas := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	alpha := make([]byte, len(digits))
	var result []string
	var dfs func(pos int)

	dfs = func(pos int) {
		if pos == len(digits) {
			result = append(result, string(alpha))
			return
		}
		for _, d := range digitsAlphas[digits[pos]] {
			alpha[pos] = d
			dfs(pos + 1)
		}
	}

	dfs(0)
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
