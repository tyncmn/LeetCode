package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Print(isPalindrome(122))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var i int64 = 1
// 	var l1_num int64 = 0
// 	var l2_num int64 = 0

// 	for l1 != nil {
// 		l1_num += int64(l1.Val) * i
// 		l1 = l1.Next
// 		i *= 10
// 	}
// 	fmt.Printf("num1 = %d\n", l1_num)
// 	i = 1
// 	for l2 != nil {
// 		l2_num += int64(l2.Val) * i
// 		l2 = l2.Next
// 		i *= 10
// 	}
// 	fmt.Printf("num2 = %d\n", l2_num)

// 	summer := l1_num + l2_num
// 	fmt.Printf("opshi cem = %d\n", summer)

// 	head := &ListNode{Val: int(summer % 10)}

// 	summer /= 10
// 	current := head

// 	for summer > 0 {
// 		newNode := &ListNode{Val: int(summer % 10)}
// 		current.Next = newNode
// 		current = newNode
// 		summer /= 10
// 	}

// 	return head
// }
