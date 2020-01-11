package linkedlists

/*
Company: Amazon(3), Microsfot(3); LinkedIn(3); Hulu(2)

Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL

Example 2:
Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
*/

func rotateRight61(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	len := 2
	tmp := head.Next
	for tmp.Next != nil {
		tmp = tmp.Next
		len++
	}
	if k == len {
		return head
	}
	tmp.Next = head
	for i := 1; i < len-(k%len); i++ {
		head = head.Next
	}
	res := head.Next
	head.Next = nil
	return res
}
