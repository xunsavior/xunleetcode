package hashtable

/*
Company: Google(2), Adobe(2); Amazon(2), Microsoft(2); Airbnb

Given an array of integers and an integer k, find out whether there are two distinct
indices i and j in the array such that nums[i] = nums[j] and the absolute difference
between i and j is at most k.

Example 1:
Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func containsNearbyDuplicate219(nums []int, k int) bool {
	dict := make(map[int]int)
	for i, v := range nums {
		if lastIndex, ok := dict[v]; ok {
			if i-lastIndex <= k {
				return true
			}
		}
		dict[v] = i
	}
	return false
}
