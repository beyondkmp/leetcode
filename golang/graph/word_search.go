package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	if len(word) > m*n {
		return false
	}

	w := []byte(word)
	var dfs func(i, j int, pos int) bool
	dfs = func(i, j int, pos int) bool {
		if pos == len(w) {
			return true
		}

		if i < 0 || j < 0 || i == m || j == n || w[pos] != board[i][j] {
			return false
		}

		board[i][j] ^= 255
		res := dfs(i, j+1, pos+1) || dfs(i, j-1, pos+1) || dfs(i+1, j, pos+1) || dfs(i-1, j, pos+1)
		board[i][j] ^= 255
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
