/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	var slow, fast *ListNode = head, head

	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	start := slow.Next
	slow.Next = reverse(start)

	ptr1 := head
	ptr2 := slow.Next
	for ptr2 != nil {
		if ptr1.Val != ptr2.Val {
			return false
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return true
}