package bs

/*
Company: Google(5), Amazon(4), Facebook(3), Uber(2); nil; Quora(3), Microsoft(3)

Let's call an array A a mountain if the following properties hold:
A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:
Input: [0,1,0]
Output: 1

Example 2:
Input: [0,2,1,0]
Output: 1

Note:
3 <= A.length <= 10000
0 <= A[i] <= 10^6
A is a mountain, as defined above.
*/

// Time: O(logN)
// Space: O(1)
func peakIndexInMountainArray852(A []int) int {
	start, end := 0, len(A)-1

	for end-start > 1 {
		mid := start + (end-start)/2
		if A[mid] < A[mid+1] {
			start = mid + 1
		} else if A[mid] > A[mid+1] {
			end = mid
		}
	}

	if A[end] > A[start] {
		return end
	}

	return start
}
