package tree

import "math"

/*
Company: LinkedIn(7), Amazon(2)
Given a non-empty binary search tree and a target value, find k values
in the BST that are closest to the target.

Note:
Given target value is a floating point.
You may assume k is always valid, that is: k ≤ total nodes.
You are guaranteed to have only one unique set of k values in the BST that are closest to the target.

Example:
Input: root = [4,2,5,1,3], target = 3.714286, and k = 2
    4
   / \
  2   5
 / \
1   3
Output: [4,3]
*/

func closestKValues272(root *TNode, target float64, k int) []int {
	output := []int{}
	helper272(root, target, k, &output)
	return output
}

func helper272(current *TNode, target float64, k int, ouput *[]int) {
	if current == nil {
		return
	}

	helper272(current.Left, target, k, ouput)

	if len(*ouput) == k {
		if math.Abs(float64(current.Val)-target) < math.Abs(float64((*ouput)[0])-target) {
			*ouput = (*ouput)[1:]
		} else {
			return
		}
	}
	*ouput = append(*ouput, current.Val)

	helper272(current.Right, target, k, ouput)
}
