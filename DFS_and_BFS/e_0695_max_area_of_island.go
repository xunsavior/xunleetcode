package dfsandbfs

/*
Company: Qualitrics(7), Amazon(5), DoorDash(5), Twitch(4), Bloomberg(3)

Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land)
connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

Example 2:
[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
*/
func maxAreaOfIsland695(grid [][]int) int {
	res := 0
	m = len(grid)
	if m == 0 {
		return res
	}
	n = len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				currentVal := helper695(grid, i, j)
				if currentVal > res {
					res = currentVal
				}
			}
		}
	}
	return res
}

// DFS
func helper695(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 { // '0'
		return 0
	}
	grid[i][j] = 0
	return 1 + helper695(grid, i+1, j) + helper695(grid, i-1, j) + helper695(grid, i, j+1) + helper695(grid, i, j-1)
}
