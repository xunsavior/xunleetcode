package linkedlists

/*
Company: Amazon(104), Microsoft(16), Adobe(11), Facebook(6), Bloomberg(6), Apple(4)
Merge two sorted linked lists and return it as a new list. The new list
should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

func mergeTwoLists21(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l2 == nil {
		return &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists21(l1.Next, nil),
		}
	}

	if l1 == nil {
		return &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists21(nil, l2.Next),
		}
	}

	if l1.Val < l2.Val {
		return &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists21(l1.Next, l2),
		}
	}

	return &ListNode{
		Val:  l2.Val,
		Next: mergeTwoLists21(l1, l2.Next),
	}
}
