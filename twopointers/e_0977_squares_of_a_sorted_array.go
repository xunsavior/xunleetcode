package twopointers

/*
Company: Facebook(11), eBay(4), Apple(4), Amazon(3), Paypal(3)

Given an array of integers A sorted in non-decreasing order, return an array
of the squares of each number, also in sorted non-decreasing order.

Example 1:
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:
1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

func sortedSquares977(A []int) []int {
	res := make([]int, len(A))
	start, end, index := 0, len(A)-1, len(A)-1
	for start <= end {
		svalue, evalue := A[start], A[end]
		if v1, v2 := svalue*svalue, evalue*evalue; v1 < v2 {
			end--
			res[index] = v2
		} else {
			start++
			res[index] = v1
		}
		index--
	}
	return res
}
