/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var ptr1, ptr2 *ListNode = list1, list2
	var head, tmp *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		head = list2
		ptr2 = ptr2.Next
	} else {
		head = list1
		ptr1 = ptr1.Next
	}
	tmp = head

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val > ptr2.Val {
			tmp.Next = ptr2
			ptr2 = ptr2.Next
		} else {
			tmp.Next = ptr1
			ptr1 = ptr1.Next
		}
		tmp = tmp.Next
	}

	if ptr1 == nil {
		tmp.Next = ptr2
	} else {
		tmp.Next = ptr1
	}

	return head
}
