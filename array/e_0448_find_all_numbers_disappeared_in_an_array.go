package array

import "math"

/*
Company: Apple(2); Google(3), Microsoft(2); Adobe(5), Facebook(2)

Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:
Input:
[4,3,2,7,8,2,3,1]
Output:
[5,6]
*/

func findDisappearedNumbers448(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if newIndex := int(math.Abs(float64(v))) - 1; nums[newIndex] > 0 {
			nums[newIndex] *= -1
		}
	}

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}
	return res
}
