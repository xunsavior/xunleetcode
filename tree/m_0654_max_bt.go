package tree

/*
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:
      6
    /   \
   3     5
    \    /
     2  0
       \
		1

Note: The size of the given array will be in the range [1,1000].
*/

func constructMaximumBinaryTree654(nums []int) *TNode {
	if len(nums) == 0 {
		return nil
	}
	i, v := helper654(nums)
	return &TNode{
		Val:   v,
		Left:  constructMaximumBinaryTree654(nums[:i]),
		Right: constructMaximumBinaryTree654(nums[i+1:]),
	}
}

// get max val of slice
func helper654(nums []int) (index int, value int) {
	if len(nums) == 1 {
		return 0, nums[0]
	}

	for i, v := range nums {
		if v > value {
			value = v
			index = i
		}
	}
	return
}
