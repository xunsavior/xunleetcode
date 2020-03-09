package tree

/*
Company: Amazon(4), Microsoft(2); ; Google(2), Snapchat(2), Facebook(1), Samsung(1)

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

	nums := []int{}
	helper653(root, &nums)

	// two-pointer algorithm
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]

		if sum == k {
			return true
		}

		if k < sum {
			j--
		} else {
			i++
		}
	}

	return false
}

// inorder traversal and save val into a slice
func helper653(root *TNode, nums *[]int) {
	if root == nil {
		return
	}

	helper653(root.Left, nums)
	val := root.Val
	*nums = append(*nums, val)
	helper653(root.Right, nums)
}
