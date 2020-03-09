package tree

/*
Company: Amazon(2)
Given a BST, convert it to a Greater Tree such that every
key of the original BST is changed to the original key plus
sum of all keys greater than the original key in BST.

Example:
Input: The root of a Binary Search Tree like this:
              5
             / \
            2   13
Output: The root of a Greater Tree like this:
             18
            /  \
          20    13
*/

func convertBST538(root *TNode) *TNode {
	if root == nil {
		return nil
	}
	lastNode := &TNode{}
	helper538(root, lastNode)
	return root
}

/*
	We use reverse-order traversal, right -> root -> left
*/
func helper538(current *TNode, lastNode *TNode) {
	if current == nil {
		return
	}

	helper538(current.Right, lastNode)

	current.Val += lastNode.Val
	lastNode.Val = current.Val

	helper538(current.Left, lastNode)

}
