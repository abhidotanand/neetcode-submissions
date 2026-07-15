/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseLL(head *ListNode) *ListNode {
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var first, last *ListNode
	var start, end *ListNode
	var cnt int

	for cnt < left - 1 {
		if first == nil {
			first = head
		} else {
			first = first.Next
		}
		cnt++
	}
	
	cnt = 0

	for cnt < right {
		if last == nil {
			last = head
		} else {
			last = last.Next
		}
		cnt++
	}

	end = last
	last = end.Next
	end.Next = nil

	if first != nil {
		start = first.Next
	} else {
		start = head
	}
	fmt.Println(first, start, end, last)
	h := reverseLL(start)

	if first == nil && last != nil{
		start.Next = last
		return h
	} else if first != nil && last == nil {
		first.Next = h
	} else if first != nil && last != nil {
		first.Next = h
		start.Next = last
	} else {
		return h
	}
	
	return head

}
