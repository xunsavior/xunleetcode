package tree

/*
Company: Facebook, Samsung

Given a Binary Search Tree and a target number,
return true if there exist two elements in the
BST such that their sum is equal to the given target.

Example 1:
Input:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 9
Output: True


Example 2:
Input:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 28
Output: False
*/

func findTarget653(root *TNode, k int) bool {
	if root == nil {
		return false
	}
	nums := map[int]int{}
	helper653(root, nums)
	for key := range nums {
		target := k - key
		if target == key {
			if nums[target] > 1 {
				return true
			}
		} else {
			if nums[target] > 0 {
				return true
			}
		}
	}
	return false
}

// inorder traversal and save val into map
func helper653(current *TNode, nums map[int]int) {
	if current == nil {
		return
	}

	if current.Left != nil {
		helper653(current.Left, nums)
	}
	// important: save val as key, and the total number of times as value
	nums[current.Val]++

	if current.Right != nil {
		helper653(current.Right, nums)
	}
}
