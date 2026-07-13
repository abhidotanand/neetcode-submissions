/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l1P, l2P *ListNode = l1, l2
    var ans, tmp *ListNode
    var carry, summ int

    for l1P != nil && l2P != nil {
        summ = (l1P.Val + l2P.Val + carry)%10
        carry = (l1P.Val + l2P.Val + carry)/10
        fmt.Println(carry, summ)
        nn := ListNode{summ, nil}
        if ans == nil {
            ans = &nn
            tmp = &nn
        } else {
            tmp.Next = &nn
            tmp = tmp.Next
        }
        l1P = l1P.Next
        l2P = l2P.Next
    }

    for l2P != nil {
        summ = (l2P.Val + carry)%10
        carry = (l2P.Val + carry)/10
        nn := ListNode{summ, nil}
        tmp.Next = &nn
        tmp = tmp.Next
        l2P = l2P.Next
    }

    for l1P != nil {
        summ = (l1P.Val + carry)%10
        carry = (l1P.Val + carry)/10
        nn := ListNode{summ, nil}
        tmp.Next = &nn
        tmp = tmp.Next
        l1P = l1P.Next
    }

    if carry > 0 {
        nn := ListNode{carry, nil}
        tmp.Next = &nn
    }

    return ans
}
