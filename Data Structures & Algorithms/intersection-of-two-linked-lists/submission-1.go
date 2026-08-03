/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int
	var tmp1, tmp2 *ListNode

	tmp1 = headA
	for tmp1 != nil {
		lenA++
		tmp1 = tmp1.Next
	}

	tmp2 = headB
	for tmp2 != nil {
		lenB++
		tmp2 = tmp2.Next
	}

	tmp1 = headA
	for lenA > lenB {
		tmp1 = tmp1.Next
		lenA--
	}

	tmp2 = headB
	for lenB > lenA {
		tmp2 = tmp2.Next
		lenB--
	}

	for tmp1 != nil && tmp2 != nil {
		if tmp1 == tmp2 {
			return tmp1
		} else if tmp1.Next == tmp2.Next {
			return tmp1.Next
		}
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
	return nil
}
