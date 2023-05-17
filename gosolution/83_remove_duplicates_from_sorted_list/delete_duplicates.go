package main

import (
	. "github.com/chinabobo/leetcode/model"
)

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		// 全部删除完再移动到下一个元素
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
