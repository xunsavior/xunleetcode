package hashtable

/*
Company: Facebook(7), Bloomberg(3), Amazon(2); Microsoft(2), Google(2)

You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water,
and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:
Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]
Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:
*/

func islandPerimeter463(grid [][]int) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if val := grid[i][j]; val == 1 {
				if j-1 < 0 || grid[i][j-1] == 0 {
					res++
				}
				if j+1 > len(grid[i])-1 || grid[i][j+1] == 0 {
					res++
				}
				if i-1 < 0 || grid[i-1][j] == 0 {
					res++
				}
				if i+1 > len(grid)-1 || grid[i+1][j] == 0 {
					res++
				}
			}
		}
	}

	return res
}
