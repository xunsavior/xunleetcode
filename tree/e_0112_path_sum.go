package tree

/*
Company: Amazon(5), Facebook(3); Microsoft(2), Adobe(2); Zillow(4)

Given a binary tree and a sum, determine if the tree has a root-to-leaf path
such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:
Given the below binary tree and sum = 22,
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

//Solution: we add current val to the previous sum
func hasPathSum112(root *TNode, sum int) bool {
	if root == nil {
		return false
	}

	l, r := root.Left, root.Right

	if root.Val == sum && l == nil && r == nil {
		return true
	}

	return hasPathSum112(l, sum-root.Val) || hasPathSum112(r, sum-root.Val)
}
