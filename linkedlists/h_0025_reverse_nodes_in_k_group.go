package linkedlists

/*
Company: Mathworks(7), Microsoft(6), Amazon(5), ByteDance(3); Facebook(5)

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number
of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:
Given this linked list: 1->2->3->4->5
For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5

Note:
Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

func reverseKGroup25(head *ListNode, k int) *ListNode {
	i := 0
	reverseHead, res := head, head
	var (
		reverseTail, lastTail *ListNode
	)
	isFirstTime := true
	for head != nil {
		i++
		if i == k {
			// perform reverse here
			reverseTail, head = head, head.Next
			reverseTail.Next = nil
			var prev *ListNode
			// remember to store the head before reversing the linked list, because it will become tail node after the reversing
			start := reverseHead
			for reverseHead != nil {
				prev, reverseHead.Next, reverseHead = reverseHead, prev, reverseHead.Next
			}
			if isFirstTime {
				res, isFirstTime = prev, false
			}
			/*
				important: we must ensure that the tail node of last reversed linked list ought
				to point to the head node of current reversed linked list
			*/
			if lastTail != nil {
				lastTail.Next = prev
			}
			// remember to update the tail node of last reversed linked list
			lastTail, start.Next, reverseHead = start, head, head
			i = 0
		} else {
			head = head.Next
		}
	}
	return res
}
