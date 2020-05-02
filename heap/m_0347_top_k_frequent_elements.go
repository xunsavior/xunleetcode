package heap

import "container/heap"

/*
Top interviewed question
Company: Amazon(22), Facebook(10), Oracle(5), Google(3), Microsoft(3)

Given a non-empty array of integers, return the k most frequent elements.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
You can return the answer in any order.
*/

type pair struct {
	val   int
	count int
}

type pairHeap []pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent347(nums []int, k int) []int {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}

	pairHeap := &pairHeap{}
	heap.Init(pairHeap)
	for k, v := range dict {
		p := pair{val: k, count: v}
		heap.Push(pairHeap, p)
	}

	res := []int{}
	for len(res) < k {
		pop := heap.Pop(pairHeap).(pair)
		res = append(res, pop.val)
	}
	return res
}
