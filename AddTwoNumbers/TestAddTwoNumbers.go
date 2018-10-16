package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getValAndMoveNext(l *ListNode) (int, *ListNode) {
	x := 0
	if l != nil {
		x = l.Val
		l = l.Next
	}
	return x, l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	x, y, carry, summary := 0, 0, 0, 0

	//fmt.Println(result.Next)
	for l1 != nil || l2 != nil {
		fmt.Println("while")
		x, l1 = getValAndMoveNext(l1)
		y, l2 = getValAndMoveNext(l2)
		fmt.Println("x:", x)
		fmt.Println("y:", y)

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

func main() {
	//     l1 := ListNode{Val:2,
	//                    &ListNode{Val:4,
	//                              &ListNode{Val:3, nil}}}
	//     l2 := ListNode{Val:5,
	//                    &ListNode{Val:6,
	//                              &ListNode{Val:4, nil}}}

	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next = &ListNode{Val: 9}

	//fmt.Printf("%v\n", l1)
	ret := addTwoNumbers(l1, l2)
	for ret != nil {
		fmt.Printf("%v", ret.Val)
		ret = ret.Next
	}
}
