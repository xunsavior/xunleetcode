package dfsandbfs

import "github.com/xunsavior/xunleetcode/linkedlists"

/*
Company: Amazon(10), Facebook(5), Lyft(3); Oracle(2), Google(2)

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two
subtrees of every node never differ by more than 1.

Example:
Given the sorted linked list: [-10,-3,0,5,9],
One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
      0
     / \
   -3   9
   /   /
 -10  5

*/

func sortedListToBST109(head *linkedlists.ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	var prev *linkedlists.ListNode
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			fast = nil
		} else {
			prev = slow
			slow, fast = slow.Next, fast.Next.Next
		}
	}

	var l, r *linkedlists.ListNode
	// remember to check if the slow node is still point to header
	if prev != nil {
		prev.Next = nil
		l = head
	} else {
		l = nil
	}
	r = slow.Next
	slow.Next = nil

	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST109(l),
		Right: sortedListToBST109(r),
	}
}
