package linkedlists

import "fmt"

/*
	You are given two non-empty linked lists representing two non-negative integers.
	The digits are stored in reverse order and each of their nodes contain a single digit.
	Add the two numbers and return it as a linked list.

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.
	Example:
	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8

	Input: (5) + (5)
	Output: 0 -> 1

	Input: (8 -> 1) + (0)
	Output: 8 -> 1
	Explanation: 342 + 465 = 807.
*/

//AddTwoNumbers ...
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		v1, v2, sum    int
		l1Next, l2Next *ListNode
	)

	// base case: both current node are nil
	if l1 == nil && l2 == nil {
		return nil
	}

	//check if current node is nil
	if v1 = 0; l1 != nil {
		v1 = l1.Val
		l1Next = l1.Next
	}

	//check if current node is nil
	if v2 = 0; l2 != nil {
		v2 = l2.Val
		l2Next = l2.Next
	}

	//check if the sum of two current nodes are >= 10
	if sum = v1 + v2; v1+v2 >= 10 {
		sum = (v1 + v2) % 10
		switch {
		// none of them have next note
		case l1Next == nil && l2Next == nil:
			return &ListNode{
				Val:  sum,
				Next: AddTwoNumbers(&ListNode{0, nil}, &ListNode{1, nil}),
			}
		// only l1 has next node
		case l1Next != nil:
			l1Next.Val++
		// only l2 has next node
		default:
			l2Next.Val++
		}
	}

	// the sum of two current nodes are < 10
	return &ListNode{
		Val:  sum,
		Next: AddTwoNumbers(l1Next, l2Next),
	}
}

// LinkedTraversal ...
func LinkedTraversal(node *ListNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d->", node.Val)
	LinkedTraversal(node.Next)
}
