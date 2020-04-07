package bs

import "sort"

/*
Company: Facebook(22), Oracle(5), Amazon(3), Microsoft(2), LinkedIn(2)
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
Each element in the result must be unique.
The result can be in any order.
*/

func intersection349(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}

	sort.Ints(nums2)
	dic := map[int]bool{}
	res := []int{}
	for _, v := range nums1 {
		if helper349(nums2, v, 0, len(nums2)-1) != -1 {
			if dic[v] == false {
				res = append(res, v)
				dic[v] = true
			}
		}
	}
	return res
}

// Binary search
func helper349(nums []int, target, low, high int) int {
	index := -1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < target {
			high = mid + 1
		} else if nums[mid] > target {
			low = mid - 1
		} else {
			index = mid
			break
		}
	}
	return index
}
