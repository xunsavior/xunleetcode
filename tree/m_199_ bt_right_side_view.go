package tree

/*
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/

func rightSideView199(root *TNode) []int {
	nums := map[int]int{}
	helper199(root, 0, nums)
	mapSize := len(nums)
	output := make([]int, 0, mapSize)
	for i := 0; i < mapSize; i++ {
		output = append(output, nums[i])
	}
	return output
}

/*
  We implement preorder traversal and store the value of each
  layer from left to right, so the right value will always overwrite
  the left value
*/
func helper199(current *TNode, depth int, nums map[int]int) {
	if current == nil {
		return
	}
	nums[depth] = current.Val
	helper199(current.Left, depth+1, nums)
	helper199(current.Right, depth+1, nums)
}
