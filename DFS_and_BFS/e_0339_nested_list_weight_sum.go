package dfsandbfs

/*
Company: LinkedIn(8), Facebook(4), Amazon(2), Cloudra(2)
Given a nested list of integers, return the sum of all integers in the list weighted by their depth.
Each element is either an integer, or a list -- whose elements may also be integers or other lists.

Example 1:
Input: [[1,1],2,[1,1]]
Output: 10
Explanation: Four 1's at depth 2, one 2 at depth 1.

Example 2:
Input: [1,[4,[6]]]
Output: 27
Explanation: One 1 at depth 1, one 4 at depth 2, and one 6 at depth 3; 1 + 4*2 + 6*3 = 27.
*/

// NestedInteger is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct{}

// IsInteger returns true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return false
}

// GetInteger return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// SetInteger set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Add this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func depthSum339(nestedList []*NestedInteger) int {
	res := 0
	if len(nestedList) == 0 {
		return res
	}

	for _, v := range nestedList {
		if v.IsInteger() {
			res += v.GetInteger()
		} else {
			helper339(v.GetList(), 2, &res)
		}
	}

	return res
}

func helper339(nestedList []*NestedInteger, depth int, res *int) {
	for _, v := range nestedList {
		if v.IsInteger() {
			*res += v.GetInteger() * depth
		} else {
			depth++
			helper339(v.GetList(), depth, res)
		}
	}
}
