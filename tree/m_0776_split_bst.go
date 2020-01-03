package tree

/*
Company: Google(2), Amazon, Coupang

Given a Binary Search Tree (BST) with root node root, and a target value V, split the tree into two
subtrees where one subtree has nodes that are all smaller or equal to the target value, while the
other subtree has all nodes that are greater than the target value.  It's not necessarily the case
that the tree contains a node with value V.

Additionally, most of the structure of the original tree should remain.  Formally, for any child C
with parent P in the original tree, if they are both in the same subtree after the split, then node C
should still have the parent P.

You should output the root TNode of both subtrees after splitting, in any order.

Example 1:
Input: root = [4,2,6,1,3,5,7], V = 2
Output: [[2,1],[4,3,6,null,null,5,7]]
Explanation:
Note that root, output[0], and output[1] are TNode objects, not arrays.

The given tree [4,2,6,1,3,5,7] is represented by the following diagram:

          4
        /   \
      2      6
     / \    / \
    1   3  5   7

while the diagrams for the outputs are:

          4
        /   \
      3      6      and    2
            / \           /
           5   7         1
Note:
The size of the BST will not exceed 50.
The BST is always valid and each node's value is different.
*/

func splitBST776(root *TNode, V int) []*TNode {
	if root == nil {
		return []*TNode{nil, nil}
	}

	// V == root.Val
	if root.Val == V {
		r := root.Right
		root.Right = nil
		return []*TNode{root, r}
	}

	l, r := helper776(root, nil, V), root
	if l != nil && r != nil {
		if l.Val == r.Val {
			return []*TNode{l, r.Right}
		}
	}
	return []*TNode{l, r}
}

func helper776(current, parent *TNode, V int) *TNode {
	if current == nil {
		return nil
	}

	if V > current.Val {
		if parent != nil {
			if parent.Left == current {
				parent.Left = current.Right
			} else {
				parent.Right = current.Right
			}
			return &TNode{
				Val:   current.Val,
				Left:  current.Left,
				Right: helper776(current.Right, parent, V),
			}
		}
		return &TNode{
			Val:   current.Val,
			Left:  current.Left,
			Right: helper776(current.Right, current, V),
		}
	}

	if V < current.Val {
		return helper776(current.Left, current, V)
	}

	if parent != nil {
		if parent.Left == current {
			parent.Left = current.Right
		} else {
			parent.Right = current.Right
		}
		current.Right = nil
	}

	return current
}
