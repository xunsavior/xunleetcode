package tree

/*
Company: Facebook, Amazon

Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

Note: A leaf is a node with no children.

Example:
Input: [1,2,3]
    1
   / \
  2   3
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Example 2:
Input: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
*/

func sumNumbers129(root *TNode) int {
	return helper129(root, 0)
}

// top-down solution
func helper129(current *TNode, prefix int) int {
	if current == nil {
		return 0
	}

	l, r := current.Left, current.Right

	if l == nil && r == nil {
		return prefix*10 + current.Val
	}

	return helper129(l, prefix*10+current.Val) + helper129(r, prefix*10+current.Val)
}
