package tree

import "strconv"

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
	sum := 0
	if root == nil {
		return sum
	}
	helper129(root, "", &sum)
	return sum
}

// top-down solution
func helper129(current *TNode, nums string, sum *int) {
	if current == nil {
		return
	}
	l, r := current.Left, current.Right

	if l == nil && r == nil {
		if num, err := strconv.Atoi(nums + strconv.Itoa(current.Val)); err == nil {
			*sum += num
		}
		return
	}

	helper129(current.Left, nums+strconv.Itoa(current.Val), sum)
	helper129(current.Right, nums+strconv.Itoa(current.Val), sum)
}
