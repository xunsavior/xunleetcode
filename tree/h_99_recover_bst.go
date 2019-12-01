package tree

/*
Two elements of a binary search tree (BST) are swapped by mistake.
Recover the tree without changing its structure.

Example 1:
Input: [1,3,null,null,2]

   1
  /
 3
  \
   2

Output: [3,1,null,null,2]

   3
  /
 1
  \
   2

Example 2:
Input: [3,1,4,null,null,2]
  3
 / \
1   4
   /
  2

Output: [2,1,4,null,null,3]
  2
 / \
1   4
   /
  3

Follow up:
A solution using O(n) space is pretty straight forward.
Could you devise a CONSTANT O(1) space solution?
*/

// Solution 1: using morris inorder traversal
func recoverTree99(root *TNode) {
	var (
		predecessor, pred, x, y *TNode
	)
	for root != nil {
		// If there is a left child, then compute the predecessor
		// If there is no link predecessor.right = root --> set it
		// If there is a linked predecessor.right = root --> break it
		if root.Left != nil {
			// Predecessor node is one step left
			// and then right till you can.
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// set the link and go to explore left subtree
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				// break link predecessor.right = root
				// link is broken : time to change subtree and go right
				// check for swap
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root

				predecessor.Right = nil
				root = root.Right
			}
		} else {
			// check for the swapped nodes
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root

			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
