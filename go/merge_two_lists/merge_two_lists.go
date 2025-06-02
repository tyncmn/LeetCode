package merge_two_lists

import listnode "leetcode/list_node"

type ListNode = listnode.ListNode

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head
	for list1 != nil {
		if list2 != nil && list1.Val > list2.Val {
			temp.Val = list1.Val
		}
	}
	return head
}
