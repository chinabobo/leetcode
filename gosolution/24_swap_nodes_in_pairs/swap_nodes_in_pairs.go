package main

import (
	. "github.com/chinabobo/leetcode/model"
)

/**
https://leetcode.cn/problems/swap-nodes-in-pairs/

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func swapPairs(head *ListNode) *ListNode {
	// 思路：将链表翻转转化为一个子问题，然后通过递归方式依次解决
	// 先翻转两个，然后将后面的节点继续这样翻转，然后将这些翻转后的节点连接起来
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next
	// 翻转当前阶段指针
	next := head.Next
	next.Next = head
	head.Next = helper(nextHead)
	return next
}
