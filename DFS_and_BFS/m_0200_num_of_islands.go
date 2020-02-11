package dfsandbfs

/*
Company: Amazon(207), Bloomberg(28), Facebook(23), Google(11), Oracle(9)
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded
by water and is formed by connecting adjacent lands horizontally or vertically.

You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input:
11110
11010
11000
00000
Output: 1

Example 2:
Input:
11000
11000
00100
00011
Output: 3
*/

var m, n int

func numIslands200(grid [][]byte) int {
	res := 0
	m = len(grid)
	if m == 0 {
		return 0
	}
	n = len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 49 { // '1'
				helper200(grid, i, j)
				res++
			}
		}
	}
	return res
}

// DFS
func helper200(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 48 { // '0'
		return
	}
	grid[i][j] = 48
	helper200(grid, i+1, j)
	helper200(grid, i-1, j)
	helper200(grid, i, j+1)
	helper200(grid, i, j-1)
}
