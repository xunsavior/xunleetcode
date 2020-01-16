package linkedlists

/*
Company: Facebook(9), Amazon(8), Microsoft(7); Cisco(3), Adobe(3)

Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may NOT modify the values in the list's nodes, only nodes itself may be changed.

Example 1:
Given 1->2->3->4, reorder it to 1->4->2->3.

Example 2:
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
*/

func reorderList143(head *ListNode) {
	// two pointer slow and fast
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			break
		} else {
			fast, slow = fast.Next.Next, slow.Next
		}
	}

	if slow == nil {
		return
	}
	// divide the original linklist into two segments from head -> slow and slow.Next->end
	start := slow.Next
	slow.Next = nil

	// reverse the segment two
	var prev *ListNode
	for start != nil {
		prev, start.Next, start = start, prev, start.Next
	}

	// add segment two to segment one
	restart := head
	for restart != nil && prev != nil {
		restart.Next, prev, prev.Next, restart = prev, prev.Next, restart.Next, restart.Next
	}
}
