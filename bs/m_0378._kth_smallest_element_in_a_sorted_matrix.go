package bs

/*
Company: Amazon(7), Facebook(5), Microsoft(3); Apple(2), Google(2)

Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
*/

func kthSmallest378(matrix [][]int, k int) int {
	N := len(matrix)
	start, end := matrix[0][0], matrix[N-1][N-1]

	for start < end {
		mid := start + (end-start)/2
		if helper378(matrix, mid) < k {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return end
}

func helper378(matrix [][]int, mid int) int {
	N := len(matrix)
	i, j, count := N-1, 0, 0
	for i >= 0 && j < N {
		if matrix[i][j] <= mid {
			count += (i + 1)
			j++
		} else {
			i--
		}
	}
	return count
}
