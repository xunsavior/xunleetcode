package hashtable

/*
Company: Amazon(6), Google(4), Yahoon(2); Apple(3), Microsoft(2)

Given an array of size n, find the majority element. The majority element
is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
Input: [3,2,3]
Output: 3

Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2

*/

// Hash table solution
// Time: O(N)
// Space: O(N)
func majorityElement169(nums []int) int {
	dict := make(map[int]int)
	max, res := len(nums)/2+len(nums)%2, 0
	for _, v := range nums {
		dict[v]++
		if count := dict[v]; count == max {
			res = v
			break
		}
	}

	return res
}

/*
Intuition of Boyer-Moore Voting Algorithm
	If we had some way of counting instances of the majority element as +1+1 and instances
	of any other element as -1−1, summing them would make it obvious that the majority
	element is indeed the majority element.

	Time: O(N)
	Space: O(1)
*/
func majorityElementOptimalSolution(nums []int) int {
	count, res := 0, 0

	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if res == v {
			count++
		} else {
			count--
		}
	}

	return res
}
