package bs

import "math"

/*
Company: Facebook(9), Amazon(5), LinkedIn(2), Microsoft(2), Yandex(2)

Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.

Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersect350(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	res := []int{}
	dict1, dict2 := make(map[int]int), make(map[int]int)
	for _, v := range nums1 {
		dict1[v]++
	}
	for _, v := range nums2 {
		dict2[v]++
	}

	for k := range dict1 {
		if dict1[k] != 0 && dict2[k] != 0 {
			times := int(math.Min(float64(dict1[k]), float64(dict2[k])))
			for i := 0; i < times; i++ {
				res = append(res, k)
			}
		}
	}

	return res
}
