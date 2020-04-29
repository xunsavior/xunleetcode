package bitmanupulation

/*
Top interviewed question
Company: Amazon(5), Microsoft(3), Bloomberg(2), Apple(2), Google(2)

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
find the one that is missing from the array.

Example 1:
Input: [3,0,1]
Output: 2

Example 2:
Input: [9,6,4,2,3,5,7,0,1]
Output: 8

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

func missingNumber268(nums []int) int {
	res := len(nums)
	for i := range nums {
		res ^= i ^ nums[i]
	}
	return res
}
