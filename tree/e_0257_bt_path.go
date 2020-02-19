package tree

import "strconv"

/*
Company: Facebook(7), Amazon(4); Microsoft(2); Yandex(2), Google(2)

Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:
Input:
   1
 /   \
2     3
 \
  5
Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/

func binaryTreePaths257(root *TNode) []string {
	output := []string{}
	if root != nil {
		helper257(root, "", &output)
	}
	return output
}

/*
	top down solution
	we append the string from root to leaf, and store the appended string
	in the output[] slice when reaching very leaf node
*/
func helper257(current *TNode, str string, output *[]string) {
	if current == nil {
		return
	}

	l, r := current.Left, current.Right
	appendStr, numString := "", strconv.Itoa(current.Val)
	if len(str) > 0 {
		appendStr = str + "->" + numString
	} else {
		appendStr = numString
	}

	if l == nil && r == nil {
		*output = append(*output, str)
		return
	}

	helper257(current.Left, appendStr, output)
	helper257(current.Right, appendStr, output)
}
