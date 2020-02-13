package dfsandbfs

/*
Company: LinkedIn(11), Atlassian(3); Google(4), Amazon(3); Pocket Gems(2)

Given a binary tree, collect a tree's nodes as if you were doing this: Collect and remove all leaves, repeat until the tree is empty.

Example:
Input: [1,2,3,4,5]

          1
         / \
        2   3
       / \
      4   5

Output: [[4,5,3],[2],[1]]
Explanation:
1. Removing the leaves [4,5,3] would result in this tree:

          1
         /
        2


2. Now removing the leaf [2] would result in this tree:

          1


3. Now removing the leaf [1] would result in the empty tree:

          []
*/

func findLeaves366(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	l, r := root.Left, root.Right
	if l == nil && r == nil {
		return [][]int{[]int{root.Val}}
	}
	leaves := []int{}
	helper366(root, nil, &leaves)
	return append(findLeaves366(root), leaves)
}

func helper366(current, parent *TreeNode, leaves *[]int) {
	l, r := current.Left, current.Right
	if l == nil && r == nil {
		*leaves = append(*leaves, current.Val)
		if parent != nil {
			if current == parent.Right {
				parent.Right = nil
			} else {
				parent.Left = nil
			}
		}
		return
	}

	if l != nil {
		helper366(current.Left, current, leaves)
	}

	if r != nil {
		helper366(current.Right, current, leaves)
	}
}
