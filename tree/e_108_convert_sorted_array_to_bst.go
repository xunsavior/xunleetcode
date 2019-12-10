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
