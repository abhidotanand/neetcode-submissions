/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var cnt, eff int
	var i *ListNode = head

	for i != nil {
		cnt++
		i = i.Next
	}

	if cnt == n{
		head = head.Next
		return head
	}

	i = head
	for i != nil {
		eff++
		if eff == cnt - n {
			break
		}
		i = i.Next
	}
	
	i.Next = i.Next.Next

	return head
}
