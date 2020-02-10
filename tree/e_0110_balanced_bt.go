package tree

import "math"

/*
Company: Amazon(5), Microsoft(3); Facebook(3), Google(3); Cisco(2)

Also see: 104

Given a binary tree, determine if it is height-balanced.
For this problem, a height-balanced binary tree is defined as:
a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

Example 1:
Given the following tree [3,9,20,null,null,15,7]:
    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:
Given the following tree [1,2,2,3,3,null,null,4,4]:
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
*/

/*
	Top down solution
	complexity: O(N*logN)
*/
func isBalanced110(root *TNode) bool {
	if root == nil {
		return true
	}
	// we get the height of root's left and right subtree
	leftH, rightH := helper110(root.Left), helper110(root.Right)
	// ensure the height difference is less than 2
	return math.Abs(float64(leftH-rightH)) <= 1 && isBalanced110(root.Left) && isBalanced110(root.Right)
}

/*
	Bottom-up solution
	Use a helper method to get the height of tree
*/
func helper110(root *TNode) int {
	if root == nil {
		return 0
	}
	l, r := helper110(root.Left)+1, helper110(root.Right)+1
	return int(math.Max(float64(l), float64(r)))
}
