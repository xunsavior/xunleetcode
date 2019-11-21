package collections

/*
	There are two sorted arrays nums1 and nums2 of size m and n respectively.
	Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
	You may assume nums1 and nums2 cannot be both empty.

	Example 1:
	nums1 = [1, 3]
	nums2 = [2]
	The median is 2.0

	Example 2:
	nums1 = [1, 2]
	nums2 = [3, 4]
	The median is (2 + 3)/2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	totalLen := len1 + len2
	combSlice := make([]int, 0, totalLen/2)
	var (
		i1 int
		i2 int
	)
	for i := 0; i <= totalLen/2; i++ {
		if i1 < len1 && i2 < len2 {
			if nums1[i1] < nums2[i2] {
				combSlice = append(combSlice, nums1[i1])
				i1++
			} else {
				combSlice = append(combSlice, nums2[i2])
				i2++
			}
		} else {
			if i1 == len1 {
				combSlice = append(combSlice, nums2[i2])
				i2++
			} else {
				combSlice = append(combSlice, nums1[i1])
				i1++
			}
		}
	}
	if totalLen%2 == 0 {
		median := (float64(combSlice[len(combSlice)-1]) + float64(combSlice[len(combSlice)-2])) / 2
		return median
	}
	return float64(combSlice[len(combSlice)-1])
}
