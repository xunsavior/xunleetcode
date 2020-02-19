package tree

import "math"

/*
Company: Facebook(8), Amazon(5); Microsoft(2); Bloomberg(3), LinkedIn(2), Google(2)
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
	currentVal := float64(root.Val)
	l, r := root.Left, root.Right

	if currentVal == target {
		return int(currentVal)
	}

	if target > currentVal {
		if r == nil {
			return int(currentVal)
		}
		rVal := float64(r.Val)
		if target >= rVal {
			return closestValue270(r, target)
		}
		childRes := closestValue270(r, target)
		gap1, gap2 := target-currentVal, math.Abs(float64(childRes)-target)
		if gap1 < gap2 {
			return int(currentVal)
		}
		return childRes
	}

	if l == nil {
		return int(currentVal)
	}
	lVal := float64(l.Val)
	if target <= lVal {
		return closestValue270(l, target)
	}
	childRes := closestValue270(l, target)
	gap1, gap2 := math.Abs(currentVal-target), math.Abs(float64(childRes)-target)
	if gap1 < gap2 {
		return int(currentVal)
	}
	return childRes
}
