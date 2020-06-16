package twopointers

/*
Company: Facebook(19), Amazon(7), Microsoft(6), LinkedIn(5), Google(3)

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.

Example:
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

Redo is also needed
*/

// time: O(m+n)
func merge88(nums1 []int, m int, nums2 []int, n int) {
	//compare from ending to starting point
	i, j, k := m-1, n-1, m+n-1
	for k >= 0 {
		if j >= 0 && i >= 0 {
			if nums1[i] >= nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else {
			if j >= 0 {
				nums1[k] = nums2[j]
				j--
			} else {
				nums1[k] = nums1[i]
				i--
			}
		}
		k--
	}
}
