package tree

/*
Company: Amazon
Given the root of a binary tree, find the maximum average value of any subtree of that tree.

Example 1:
Input: [5,6,1]
Output: 6.00000

Explanation:
For the node with value = 5 we have an average of (5 + 6 + 1) / 3 = 4.
For the node with value = 6 we have an average of 6 / 1 = 6.
For the node with value = 1 we have an average of 1 / 1 = 1.
So the answer is 6 which is the maximum.

Note:
The number of nodes in the tree is between 1 and 5000.
Each node will have a value between 0 and 100000.
Answers will be accepted as correct if they are within 10^-5 of the correct answer.
*/

func maximumAverageSubtree1120(root *TNode) float64 {
	max := float64(0)
	helper1120(root, &max)
	return max
}

/*
	Bottom-up solution
	We will return TWO values to upper layer:
	1. total size of left and right subtree
	2. sum of left.Val and right.Val
*/
func helper1120(current *TNode, max *float64) (size int, sum float64) {
	if current == nil {
		size, sum = 0, float64(0)
		return
	}

	l, r := current.Left, current.Right
	lsize, lsum := helper1120(l, max)
	rsize, rsum := helper1120(r, max)

	size, sum = lsize+rsize+1, lsum+rsum+float64(current.Val)

	average := sum / float64(size)
	if average > *max {
		*max = average
	}

	return size, sum
}
