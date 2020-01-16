package linkedlists

/*
Company: Amazon(6), Microsof(3); ; Salesforce(6), Uber(2)
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
Input: 1->1->2
Output: 1->2

Example 2:
Input: 1->1->2->3->3
Output: 1->2->3
*/

func deleteDuplicates83(head *ListNode) *ListNode {
	start := head
	for start != nil {
		next := start.Next
		if next == nil {
			break
		}
		if start.Val == next.Val {
			start.Next = next.Next
		} else {
			start = start.Next
		}
	}
	return head
}
