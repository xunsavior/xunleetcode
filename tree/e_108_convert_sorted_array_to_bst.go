package tree

/*
Also see: 654, 105, 106
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of
the two subtrees of every node never differ by more than 1.

Example:
Given the sorted array: [-10,-3,0,5,9],
One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
      0
     / \
   -3   9
   /   /
 -10  5
*/
func sortedArrayToBST108(nums []int) *TNode {
	size := len(nums)
	if size == 0 {
		return nil
	}
	mid := size / 2
	return &TNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST108(nums[:mid]),
		Right: sortedArrayToBST108(nums[mid+1:]),
	}
}
