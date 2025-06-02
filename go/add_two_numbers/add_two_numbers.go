package main

import (
	listnode "leetcode/list_node"
)

type ListNode = listnode.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c int = 0
	head := &ListNode{}
	temp := head
	l1_num := 0
	l2_num := 0
	for l1 != nil || l2 != nil {

		if l1 == nil {
			l1_num = 0
		} else {
			l1_num = l1.Val
		}

		if l2 == nil {
			l2_num = 0
		} else {
			l2_num = l2.Val
		}
		a := l1_num + l2_num
		temp.Val = (a + c) % 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
		c = (a + c) / 10
		if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
			temp = temp.Next
		} else if c != 0 {
			temp.Next = &ListNode{Val: c}
		}
	}

	return head
}
