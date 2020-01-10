package twopointers

/*
Company: Facebook(10), Amazon(5), Microsoft(3), Apple(2), Google(2)
Given an array nums, write a function to move all 0's to the end
of it while maintaining the relative order of the non-zero elements.

Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:
You must do this IN-PLACE without making a copy of the array.
Minimize the total number of operations.
*/

func moveZeroes283(nums []int) {
	n := len(nums)
	start := 0
	for _, v := range nums {
		if v != 0 {
			nums[start] = v
			start++
		}
	}

	for start < n {
		nums[start] = 0
		start++
	}
}
