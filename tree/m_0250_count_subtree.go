package tree

/*
Company: Box, Bloomberg

Given a binary tree, count the number of uni-value subtrees.
A Uni-value subtree means all nodes of the subtree have the same value.

Example :
Input:  root = [5,1,5,5,5,null,5]

              5
             / \
            1   5
           / \   \
          5   5   5
Output: 4
*/

func countUnivalSubtrees250(root *TNode) int {
	if root == nil {
		return 0
	}
	count := 0
	helper250(root, &count)
	return count
}

// bottom-up solution
func helper250(current *TNode, count *int) bool {
	if current == nil {
		return true
	}
	l, r := current.Left, current.Right
	lv, rv := helper250(current.Left, count), helper250(current.Right, count)
	if lv && rv {
		if (l != nil && l.Val != current.Val) || (r != nil && r.Val != current.Val) {
			return false
		}
		*count++
		return true
	}
	return false
}
