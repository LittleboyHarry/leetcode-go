package main

import . "leetcode-go/common/listnode"

var zeroNode = &ListNode{Val: 0}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummy   = &ListNode{Val: 0}
		prev    = dummy
		p1      = l1
		p2      = l2
		carrier = 0
	)
	for p1 != p2 {
		prev.Next = &ListNode{Val: (p1.Val + p2.Val + carrier) % 10}
		prev = prev.Next
		if (p1.Val + p2.Val + carrier) >= 10 {
			carrier = 1
		} else {
			carrier = 0
		}
		if p1 = p1.Next; p1 == nil {
			p1 = zeroNode
		}
		if p2 = p2.Next; p2 == nil {
			p2 = zeroNode
		}
	}
	if carrier > 0 {
		prev.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
