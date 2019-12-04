package tree

/*
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.
Example:
	 7
	/ \
   3   15
	   / \
	  9  20

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false


Note:
next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

*/

// BSTIterator ...
type BSTIterator struct {
	Stack []*TNode
}

// Constructor ...
func Constructor(root *TNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.Stack = []*TNode{}
	helper173(root, &iterator.Stack)
	return iterator
}

/*
	ONLY store the nodes from root to the most left leaf node to stack
	To achieve O(h)
*/
func helper173(current *TNode, stack *[]*TNode) {
	if current == nil {
		return
	}
	*stack = append(*stack, current)
	helper173(current.Left, stack)
}

// Next ...
func (iterator *BSTIterator) Next() int {
	size := len(iterator.Stack)
	// get the last element because you store from root to the most leaf node
	node := iterator.Stack[size-1]
	// removed the last element to achieve pop()
	iterator.Stack = iterator.Stack[:size-1]
	// start from its right node
	helper173(node.Right, &iterator.Stack)
	return node.Val
}

// HasNext ...
func (iterator *BSTIterator) HasNext() bool {
	return len(iterator.Stack) != 0
}
