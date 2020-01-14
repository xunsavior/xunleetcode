package linkedlists

/*
Company: Microsoft(6), Amazon(3), IXL(2), Adobe(2)
Given a singly linked list, determine if it is a palindrome.

Example 1:
Input: 1->2
Output: false

Example 2:
Input: 1->2->2->1
Output: true

Follow up:
Could you do it in O(n) time and O(1) space?
*/

func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := helper234(head.Next, head.Next.Next)

	// reverse the mid node
	var prev *ListNode
	for mid != nil {
		prev, mid.Next, mid = mid, prev, mid.Next
	}

	// compare reversed head to head
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev, head = prev.Next, head.Next
	}

	return true
}

// find the mid node via two pointers
func helper234(slow *ListNode, fast *ListNode) *ListNode {
	if fast == nil {
		return slow
	}

	if fast.Next == nil {
		return slow.Next
	}

	return helper234(slow.Next, fast.Next.Next)
}
