package tree

/*
Given a binary tree, flatten it to a linked list IN-PLACE.

For example, given the following tree:
    1
   / \
  2   5
 / \   \
3   4   6

The flattened tree should look like:
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

/*
   connect current node's right subtree to the most right leaf node of its left subtree
*/
func flatten114(root *TNode) {
	if root == nil {
		return
	}
	l, r := root.Left, root.Right
	mostR := helper114(l)
	if l != nil {
		mostR.Right = r
		root.Right = l
		root.Left = nil
	}
	flatten114(root.Left)
	flatten114(root.Right)
}

// get the MOST RIGHT leaf node
func helper114(current *TNode) *TNode {
	if current == nil {
		return nil
	}
	if current.Right == nil {
		return current
	}
	return helper114(current.Right)
}
