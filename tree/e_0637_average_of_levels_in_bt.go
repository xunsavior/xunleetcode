package tree

/*
Company: Facebook(4), Amazon(2)
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]

Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

Note:
The range of node's value is in the range of 32-bit signed integer.
*/

// BFS
func averageOfLevels637(root *TNode) []float64 {
	if root == nil {
		return nil
	}
	res := []float64{}
	queue := []*TNode{root}

	for len(queue) != 0 {
		sum := 0
		nextLevel := []*TNode{}
		for _, node := range queue {
			sum += node.Val
			l, r := node.Left, node.Right
			if l != nil {
				nextLevel = append(nextLevel, l)
			}
			if r != nil {
				nextLevel = append(nextLevel, r)
			}
		}
		res = append(res, float64(sum)/float64(len(queue)))
		queue = nextLevel
	}

	return res
}
