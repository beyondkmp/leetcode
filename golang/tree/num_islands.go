package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	var dfs func(i, j int)
	var count int

	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}
