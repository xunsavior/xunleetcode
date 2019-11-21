package tree

/*
	Given two binary trees, write a function to check if they are the same or not.
	Two binary trees are considered the same if they are structurally identical
	and the nodes have the same value.

	Example 1:
	Input:     1         1
			  / \       / \
			 2   3     2   3
			[1,2,3],   [1,2,3]
	Output: true

	Example 2:
	Input:     1         1
			  /           \
			 2             2
			[1,2],     [1,null,2]

	Output: false

	Example 3:
	Input:   1         1
			/ \       / \
		   2   1     1   2
		  [1,2,1],  [1,1,2]
	Output: false
*/

// IsSameTree ...
func IsSameTree(p *TNode, q *TNode) bool {
	// check if reach the leaf node for both tree
	if p == nil && q == nil {
		return true
	}
	// check if one of the two current nodes is nil
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	//check if both node return the same value
	if p.Val != q.Val {
		return false
	}
	//continue to checking their the left and right node
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
