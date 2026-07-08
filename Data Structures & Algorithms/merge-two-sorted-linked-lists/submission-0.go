/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ptr1, ptr2 *ListNode = list1, list2
	var tmp *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val > ptr2.Val {
			tmp = ptr2.Next
			ptr2.Next = ptr1
			ptr2 = tmp
		} else {
			tmp := ptr1.Next
			ptr1.Next = ptr2
			ptr1 = tmp
		}
	}
	if ptr1 == nil {
		return list1
	}
	return list2
}
