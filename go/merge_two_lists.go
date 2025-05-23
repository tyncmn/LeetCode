package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head
	for list1 != nil {
		if list2 != nil && list1.Val > list2.Val {
			temp.Val = list1.Val
		}
	}
	return head
}
