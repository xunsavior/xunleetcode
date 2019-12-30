package tree

/*
Company: Microsoft(2), Yahoon(2), Adobe(2)

Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

Example:
Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]

Explanation:
The above output corresponds to the 5 unique BST's shown below:
   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

func generateTrees95(n int) []*TNode {
	if n == 0 {
		return nil
	}
	return helper95(1, n)
}

func helper95(start, end int) []*TNode {
	output := []*TNode{}

	if start > end {
		output = append(output, nil)
		return output
	}

	for i := start; i <= end; i++ {
		l, r := helper95(start, i-1), helper95(i+1, end)
		// form all left and right subtree cases when i is root node
		for _, lv := range l {
			for _, rv := range r {
				root := &TNode{
					Val:   i,
					Left:  lv,
					Right: rv,
				}
				output = append(output, root)
			}
		}
	}

	return output
}
