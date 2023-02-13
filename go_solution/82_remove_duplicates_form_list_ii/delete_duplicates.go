package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现的数字。
// 思路：链表头结点可能被删除，所以用 dummy node 辅助删除

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			// 记录已经删除的值，用于后续节点判断
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
