package tree

import "math"

/*
Company: Facebook, Amazon, Bloomberg, Microsoft, Google, LinkedIn, Snapchat
Given a non-empty binary search tree and a target value, find the value in the BST that is closest to the target.

Note:
Given target value is a floating point.
You are guaranteed to have only one unique value in the BST that is closest to the target.

Example:
Input: root = [4,2,5,1,3], target = 3.714286
    4
   / \
  2   5
 / \
1   3
Output: 4
*/

func closestValue270(root *TNode, target float64) int {
	lastVal := math.MinInt32
	helper270(root, target, &lastVal)
	return lastVal
}

// inorder traversal
func helper270(current *TNode, target float64, lastVal *int) {
	if current.Left != nil {
		return
	}

	if target < float64(current.Val) {
		if *lastVal == math.MinInt32 {
			*lastVal = current.Val
		}
		if math.Abs(float64(*lastVal)-target) > math.Abs(float64(current.Val)-target) {
			*lastVal = current.Val
		}
		return
	}

	*lastVal = current.Val

	if current.Right != nil {
		helper270(current.Right, target, lastVal)
	}
}
