package sorting

import "sort"

/*
Company: Facebook(25), Google(11), Microsoft(10), Amazon(9), Uber(8)
Given a collection of intervals, merge all overlapping intervals.

Example 1:
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
*/

func merge56(intervals [][]int) [][]int {
	// we need to sort based on the first val of each interval group
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for i, v := range intervals {
		// if there is no added interval  or 1st val of current interval > 2nd val of last added interval
		if i == 0 || v[0] > res[len(res)-1][1] {
			res = append(res, v)
			continue
		}
		pop := res[len(res)-1]
		// case 1: 1st val of current interval equal to 2nd val of last added interval
		// case 2: 2nd val of current interval >= 2nd val of last added interval
		if v[0] == pop[1] || v[1] > pop[1] {
			pop[1] = v[1]
		}
	}
	return res
}
