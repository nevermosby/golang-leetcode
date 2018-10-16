package addtwonumbers

/**
 * Definition for singly-linked list.
  * type ListNode struct {
  *     Val int
  *     Next *ListNode
  * }
  */:
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	x, y, carry, summary := 0, 0, 0, 0

	for l1 != nil || l2 != nil {

		x, l1 = getValAndMoveNext(l1)
		y, l2 = getValAndMoveNext(l2)

		summary = carry + x + y
		current.Next = &ListNode{Val: summary % 10}
		current = current.Next
		carry = summary / 10
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10}
	}

	return result.Next
}

func getValAndMoveNext(l *ListNode) (int, *ListNode) {
	x := 0
	if l != nil {
		x = l.Val
		l = l.Next
	}
	return x, l
}
