/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	var slow, fast *ListNode = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		} else {
			fast = fast.Next.Next
		}
	}

	curr := slow.Next
	slow.Next = nil

	var prev *ListNode
	for  curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	prev_tmp := prev
	head_tmp := head
	for prev_tmp != nil {
		prev_next := prev_tmp.Next
		head_next := head_tmp.Next

		head_tmp.Next = prev_tmp
		prev_tmp.Next = head_next
		
		head_tmp = head_next
		prev_tmp = prev_next
	}
}