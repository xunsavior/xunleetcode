package bs

/*
Company: Amazon(61), Microsoft(8), Facebook(5), Google(2), Walmart Labs(2)

Write an efficient algorithm that searches for a value in an m x n matrix.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Example:
Consider the following matrix:
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

Given target = 5, return true.
Given target = 20, return false.
*/

func searchMatrix240(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for _, row := range matrix {
		if row[0] > target {
			continue
		}
		if helper240(row, target) {
			return true
		}
	}

	return false
}

func helper240(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}

// search space reduction
func searchMatrix240SearchSpaceReduction(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix)-1, 0
	for m >= 0 && n < len(matrix[0]) {
		val := matrix[m][n]
		if target > val {
			m--
		} else if target < val {
			n++
		} else {
			return true
		}
	}

	return false
}
