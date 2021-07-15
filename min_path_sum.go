package main

// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	pathGrid := make([][]int, h)
	for i := 0; i < h; i++ {
		pathGrid[i] = make([]int, w)
	}

	pathGrid[0][0] = grid[0][0]

	for i := 1; i < h; i++ {
		pathGrid[i][0] = pathGrid[i-1][0] + grid[i][0]
	}

	for j := 1; j < w; j++ {
		pathGrid[0][j] = pathGrid[0][j-1] + grid[0][j]
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			v := grid[i][j]
			pathGrid[i][j] = min(v+pathGrid[i-1][j], v+pathGrid[i][j-1])
		}
	}

	return pathGrid[h-1][w-1]
}
