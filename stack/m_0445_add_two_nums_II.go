package stack

import "github.com/xunsavior/go-basic-utils/data_structure/linkedlist"

/*
Company: Microsoft(9), Amazon(7), Bloomberg(4), Oracle(2); Facebook(2)

You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is NOT allowed.

Example:
Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
*/

func addTwoNumbers445(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	// store l1 and l2 to stack1 and stack2
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var res *linkedlist.ListNode
	carry := 0
	for len(stack1) != 0 || len(stack2) != 0 {
		// if the stack is empty, we directly set the val to 0
		n1, n2 := 0, 0
		if len(stack1) != 0 {
			n1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			n2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		if res == nil {
			res = new(linkedlist.ListNode)
			res.Val = sum % 10
		} else {
			head := new(linkedlist.ListNode)
			head.Val, head.Next = sum%10, res
			res = head
		}
		// boundary case handling, for instance, [5], [5]
		if carry == 1 && len(stack1) == 0 && len(stack2) == 0 {
			head := new(linkedlist.ListNode)
			head.Val, head.Next = 1, res
			res = head
		}
	}

	return res
}
