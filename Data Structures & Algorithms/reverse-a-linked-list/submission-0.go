/*
Definition for singly-linked list.
*/
// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else if head.Next.Next == nil {
		head.Next.Next = head
		return head.Next
	}

	var front *ListNode = head
	var mid *ListNode = front.Next
	var rear *ListNode = mid.Next

	front.Next = nil

	for rear.Next != nil {
		mid.Next = front

		front = mid
		mid = rear
		rear = rear.Next
	}

	rear.Next = mid
	mid.Next = front

	return rear
}
