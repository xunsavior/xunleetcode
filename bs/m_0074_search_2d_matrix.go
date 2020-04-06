package bs

/*
Company: Amazon(13), Microsoft(8), Facebook(5), Uber(2), Apple(2)

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.

Example 1:
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true

Example 2:
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/

func searchMatrix74(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	top, bottom, targetRowIndex := 0, m-1, -1
	for top <= bottom {
		mid := top + (bottom-top)/2
		min, max := matrix[mid][0], matrix[mid][n-1]
		if target < min {
			bottom = mid - 1
		} else if target > max {
			top = mid + 1
		} else {
			targetRowIndex = mid
			break
		}
	}

	if targetRowIndex == -1 {
		return false
	}

	nums := matrix[targetRowIndex]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
