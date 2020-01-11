package linkedlists

/*
Company: Amazon(5), Microsoft(4), Apple(3), Facebook(2), Oracle(2)

Given a linked list, remove the n-th node from the end of list and return its head.

Example:
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.

Follow up:
Could you do this in one pass?
*/

// Two pointers
func removeNthFromEnd19(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	next := slow.Next.Next
	slow.Next = next
	slow.Next.Next = nil
	return head
}
