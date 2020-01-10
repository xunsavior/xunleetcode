package bitmanupulation

/*
Company: Apple(6), Amazon(5), Google(4), Tecent(2), Facebook(2)
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,1]
Output: 1

Example 2:
Input: [4,1,2,1,2]
Output: 4
*/

func singleNumber136(nums []int) int {
	res := 0

	for _, v := range nums {
		res ^= v
	}

	return res
}
