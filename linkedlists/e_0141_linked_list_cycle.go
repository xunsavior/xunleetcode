package linkedlists

/*
Company: Amazon(4), Apple(2), Microsoft(2), Google(2), Oracle(2)
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use
an integer pos which represents the position (0-indexed)
in the linked list where tail connects to. If pos is -1,
then there is no cycle in the linked list.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Follow up:
Can you solve it using O(1) (i.e. constant) memory?
*/

// Two pointers
func hasCycle141(head *ListNode) bool {
	if head == nil {
		return false
	}
	next := head.Next
	if next == nil {
		return false
	}
	return helper141(head, next.Next)
}

func helper141(node1 *ListNode, node2 *ListNode) bool {
	if node2 == nil {
		return false
	}
	if node1.Val == node2.Val {
		return true
	}
	if node2.Next == nil {
		return false
	}
	return helper141(node1.Next, node2.Next.Next)
}
