package linkedlists

/*
Company: Pure Storage(6), Apple(3), Bloomberg(2); Amazon(4), Google(2)
Remove all elements from a linked list of integers that have value val.

Example:
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/

func removeElements203(head *ListNode, val int) *ListNode {
	var last *ListNode
	tmp := head
	for tmp != nil {
		if tmp.Val == val {
			if last != nil {
				last.Next, tmp = tmp.Next, tmp.Next
			} else {
				head, tmp = head.Next, head.Next
			}
			continue
		}
		last, tmp = tmp, tmp.Next
	}
	return head
}
