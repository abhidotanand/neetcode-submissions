/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	if head.Next.Next == nil && head.Next.Val == val {
		head.Next = nil
		return head
	}

	tmp := head

	for tmp != nil && tmp.Next != nil {
		
		for tmp.Next != nil && tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return head
}
