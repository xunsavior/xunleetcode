package linkedlists

/*
Company: Microsoft(5), Amazon(3), Bloomberg(2), Oracle(2)
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

func reverseBetween92(head *ListNode, m int, n int) *ListNode {
	start, i := 1, 0
	if m > 1 {
		start = m - 1
	}
	var (
		startNode, endNode, reversedStart, prev *ListNode
	)
	tmp := head
	for tmp != nil {
		i++
		if i == start {
			startNode = tmp
		}
		if i == m {
			reversedStart = tmp
		}
		if i == n {
			endNode = tmp.Next
			tmp.Next = nil
		}
		tmp = tmp.Next
	}
	end := reversedStart
	for reversedStart != nil {
		prev, reversedStart.Next, reversedStart = reversedStart, prev, reversedStart.Next
	}
	startNode.Next, end.Next = prev, endNode
	return head
}
