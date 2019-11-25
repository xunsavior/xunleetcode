package tree

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

func helper108(start, end int, nums []int) *TNode {
	// base case: the number of nodes of subtree is 2
	if end-start == 1 {
		return &TNode{
			Val: nums[end],
			Left: &TNode{
				Val:   nums[start],
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}
	}

	// base case: the number of nodes of subtree is 1
	if end == start {
		return &TNode{
			Val:   nums[start],
			Left:  nil,
			Right: nil,
		}
	}

	// we split the tree into left and right subtree using mid index (i.e. (start+end)/2 )
	m := (end + start) / 2
	node := &TNode{
		Val:   nums[m],
		Left:  nil,
		Right: nil,
	}

	//get all left subtrees
	node.Left = helper108(start, m-1, nums)
	//get all right subtrees
	node.Right = helper108(m+1, end, nums)

	return node
}
