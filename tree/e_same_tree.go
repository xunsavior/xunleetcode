package tree

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
