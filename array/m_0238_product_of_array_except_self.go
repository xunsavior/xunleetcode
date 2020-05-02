package array

/*
Company: Facebook(75), Amazon(23), Lyft(22), Goldman Sachs(13), Microsoft(10)

Given an array nums of n integers where n > 1,  return an array output such that output[i]
is equal to the product of all the elements of nums except nums[i].

Example:
Input:  [1,2,3,4]
Output: [24,12,8,6]
Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.
Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
*/

func productExceptSelf238(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1

	// the product of left elements for current element
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	// the product of left elements times the right ones
	product := 1
	for i := n - 2; i >= 0; i-- {
		product = product * nums[i+1]
		res[i] = res[i] * product
	}

	return res
}
