package main

import (
	. "github.com/chinabobo/leetcode/model"
)

/*
将单链表按值划分为左边小，中间相等，右边大的形式
*/

func partition2(head *ListNode, x int) *ListNode {
	if head.Next == nil {
		return nil
	}
	headDummy := &ListNode{Val: 0}
	equalDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	headDummy.Next = head
	head = headDummy

	equal := equalDummy
	tail := tailDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else if head.Next.Val == x {
			t := head.Next
			head.Next = head.Next.Next
			equal.Next = t
			equal = equal.Next
		} else if head.Next.Val > x {
			t := head.Next
			head.Next = head.Next.Next
			tail.Next = t
			tail = tail.Next
		}
	}

	tail.Next = nil
	head.Next = equalDummy.Next
	equal.Next = tailDummy.Next
	return headDummy.Next
}
