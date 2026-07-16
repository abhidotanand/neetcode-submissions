/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
    var i *Node = head

	for i != nil {
		tmp := &Node{}

		tmp.Val = i.Val
		tmp.Next = i.Next
		i.Next = tmp

		i = tmp.Next
	}

	i = head.Next
	r := head.Random
	for i != nil {
		if r != nil {
			i.Random = r.Next
		} else {
			i.Random = nil
		}

		if i.Next != nil {
			r = i.Next.Random
			i = i.Next.Next
		} else {
			break
		}
	}

	newHead := head.Next
	i = head

	for i.Next != nil {
		tmp := i.Next
		tmp2 := i.Next.Next
		i.Next = tmp2
		i = tmp
	}

	return newHead
}
