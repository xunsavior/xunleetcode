package stack

/*
Company: Amazon(6), Facebook(2), Twitter(2); Bloomberg(2); Google(3), Mathworks(2)
You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements
are subset of nums2. Find all the next greater numbers for nums1's elements in the
corresponding places of nums2.

The Next Greater Number of a number x in nums1 is the first greater number to its
right in nums2. If it does not exist, output -1 for this number.

Example 1:
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
	For number 2 in the first array, there is no next greater number for it in the second array, so output -1.

Example 2:
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
	For number 4 in the first array, there is no next greater number for it in the second array, so output -1.

Note:
All elements in nums1 and nums2 are unique.
The length of both nums1 and nums2 would not exceed 1000.
*/

func nextGreaterElement496(nums1 []int, nums2 []int) []int {
	dict, stack, res := make(map[int]int), make([]int, 0), make([]int, 0)
	for _, v := range nums2 {
		// remember to use while loop to add previous nums if it is in descending sequence
		for len(stack) != 0 && v > stack[len(stack)-1] {
			dict[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for _, v := range nums1 {
		if num, ok := dict[v]; ok {
			res = append(res, num)
			continue
		}
		res = append(res, -1)
	}
	return res
}
