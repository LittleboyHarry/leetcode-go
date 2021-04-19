package singlylist

import . "leetcode-go/common/listnode"

func Make(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	} else {
		head := &ListNode{Val: values[0]}
		pointer := head
		for _, value := range values[1:] {
			pointer.Next = &ListNode{Val: value}
			pointer = pointer.Next
		}
		return head
	}
}

func ToSlice(head *ListNode) []int {
	var t []int
	for node := head; node != nil; node = node.Next {
		t = append(t, node.Val)
	}
	return t
}
