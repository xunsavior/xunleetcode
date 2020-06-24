package linkedlists

/*
Company: Microsoft(2), Apple(2), Adobe(2); Amazon(2); Google(3), LinkedIn(2)

Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position
(0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

*/

func detectCycle142(head *ListNode) *ListNode {
	slow, fast, thirdPointer := head, head, head

	for slow != nil {
		slow = slow.Next

		if slow == nil {
			break
		}

		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			break
		}

		fast = fast.Next.Next
		// first encounter
		if slow == fast {
			// so the third pointer start moving from head
			for {
				// if the slow pointer encounters with the third pointer, then return the res
				if slow == thirdPointer {
					return slow
				}
				slow = slow.Next
				thirdPointer = thirdPointer.Next
			}
		}

	}

	return nil
}
