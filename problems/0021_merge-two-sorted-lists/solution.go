package _021_merge_two_sorted_lists

import . "leetcode-go/common/listnode"

type context struct {
	p1 *ListNode
	p2 *ListNode
}

const (
	dropEmpty = iota
	dropLeft
	dropRight
)

func (ctx *context) pop() *ListNode {
	e1, e2 := ctx.p1 != nil, ctx.p2 != nil
	var policy int
	if !e1 && !e2 {
		policy = dropEmpty
	} else if !e1 && e2 {
		policy = dropRight
	} else if e1 && !e2 {
		policy = dropLeft
	} else {
		if ctx.p1.Val < ctx.p2.Val {
			policy = dropLeft
		} else {
			policy = dropRight
		}
	}
	switch policy {
	case dropLeft:
		r1 := ctx.p1
		ctx.p1 = ctx.p1.Next
		return r1
	case dropRight:
		r2 := ctx.p2
		ctx.p2 = ctx.p2.Next
		return r2
	default:
		return nil
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	c := context{l1, l2}
	head := c.pop()
	curr := head
	for next := c.pop(); next != nil; next = c.pop() {
		curr.Next = next
		curr = next
	}
	return head
}
