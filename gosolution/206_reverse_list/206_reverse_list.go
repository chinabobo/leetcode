package main

import (
	. "github.com/chinabobo/leetcode/model"
)

// 反转一个单链表。
// 思路：用一个 prev 节点保存向前指针，temp 保存向后的临时指针

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
