/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverser(head *ListNode) *ListNode {
	var prev *ListNode
	var curr *ListNode = head
	var tmp *ListNode
	
	for curr != nil {
		tmp = curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func pairSum(head *ListNode) int {
	var ans int
    var mid, fast, tmp1 *ListNode = head, head, head

	for fast != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}

	tmp2 := reverser(mid)

	for tmp2 != nil {
		ans = max(ans, tmp1.Val + tmp2.Val)
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}

	return ans
}