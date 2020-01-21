package linkedlists

/*
Company: Google(3); Amazon(2)
Given a non-negative integer represented as non-empty a singly linked list of digits, plus one to the integer.
You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.

Example :
Input: [1,2,3]
Output: [1,2,4]
*/

func plusOne369(head *ListNode) *ListNode {
	var (
		prev, newHead *ListNode
	)
	// reverse
	for head != nil {
		prev, head.Next, head = head, prev, head.Next
	}
	// iteration
	start, isCarryNeeded := prev, true
	for start != nil && isCarryNeeded {
		if start.Val+1 > 9 {
			start.Val = 0
			if start.Next == nil {
				start.Next = new(ListNode)
				start.Next.Val = 1
				isCarryNeeded = false
			}
		} else {
			start.Val++
			isCarryNeeded = false
		}
		start = start.Next
	}
	// reverse back again
	for prev != nil {
		newHead, prev.Next, prev = prev, newHead, prev.Next
	}

	return newHead
}
