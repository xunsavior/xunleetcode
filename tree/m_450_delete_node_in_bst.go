package tree

/*
Company: Microsoft(3), Oracle(2), Amazon(3), Google(2), Facebook(2)

Given a root node reference of a BST and a key, delete the node with the given key in the BST.
Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:
Search for a node to remove.
If the node is found, delete the node.

Note: Time complexity should be O(height of tree).

root = [5,3,6,2,4,null,7]
key = 3
    5
   / \
  3   6
 / \   \
2   4   7
Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
*/

func deleteNode450(root *TNode, key int) *TNode {
	if root.Left == nil || root.Right == nil && root.Val == key {
		return nil
	}
	helper450(root, nil, key)
	return root
}

func helper450(current, parent *TNode, key int) {
	if current == nil {
		return
	}
	if key > current.Val {
		helper450(current.Right, current, key)
	} else if key < current.Val {
		helper450(current.Left, current, key)
	} else {
		l, r := current.Left, current.Right
		if l == nil && r == nil {
			if parent.Left == current {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			return
		}
		if l != nil {
			leafRight := searchLeaf(l, false)
			current.Val, leafRight.Val = leafRight.Val, current.Val
			helper450(l, current, key)
			return
		}
		if r != nil {
			leafLeft := searchLeaf(r, true)
			current.Val, leafLeft.Val = leafLeft.Val, current.Val
			helper450(r, current, key)
			return
		}
	}
}

func searchLeaf(current *TNode, isLeft bool) *TNode {
	if isLeft {
		if current.Left == nil {
			return current
		}
		return searchLeaf(current.Left, isLeft)
	}

	if current.Right == nil {
		return current
	}

	return searchLeaf(current.Right, isLeft)
}
