package twopointers

/*
Company: Amazon(10), Microsoft(7), Oracle(7), Apple(5), LinkedIn(4)

Given an array with n objects colored red, white or blue, sort them in-place so that objects
of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Follow up:
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
*/

// one-pass and in-place
func sortColors75(nums []int) {
	start, end := 0, len(nums)-1
	var curr int
	for curr <= end {
		if val := nums[curr]; val == 0 {
			// swap the value between position start and curr if current val is 0
			nums[curr], nums[start] = nums[start], val
			start++
			curr++
		} else if val == 2 {
			// swap the value between position curr and end if current val is 2
			nums[curr], nums[end] = nums[end], val
			end--
		} else {
			curr++
		}
	}
}
