package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
	"github.com/chinabobo/leetcode/util/linkedlist"
)

// https://leetcode.cn/problems/partition-list/
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 思路：将大于x的节点，放到另外一个链表，最后连接这两个链表

func main() {
	list1 := linkedlist.NewLinkedList([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(linkedlist.PrintLinkedList(partition(list1, 3)))
}

func partition(head *ListNode, x int) *ListNode {
	// check
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	headDummy.Next = head
	head = headDummy

	tail := tailDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除<x节点
			t := head.Next
			head.Next = head.Next.Next
			// 放到另外一个链表
			tail.Next = t
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
