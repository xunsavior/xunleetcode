package linkedlists

/*
Important
Company: Amazon(19), Microsoft(16), Apple(9), Mathworks(9), Google(3)
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Follow up:
A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

// dummy -> 1 -> nil
// dummy -> 2 -> 1 -> nil
// dummy -> 3 -> 2 -> 1 -> nil
func reverseList206(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head.Next, head = head, prev, head.Next
	}
	return prev
}
