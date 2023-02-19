package main

import "fmt"

// https://leetcode.cn/problems/partition-list/
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 思路：将大于x的节点，放到另外一个链表，最后连接这两个链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(arr []int) *ListNode {
	var head ListNode
	var pre ListNode
	for _, num := range arr {
		node := ListNode{Val: num, Next: nil}
		if head.Next == nil {
			head.Next = &node
		}
		if pre.Next == nil {
			pre.Next = &node
		} else {
			pre.Next.Next = &node
			pre = *pre.Next
		}
	}
	return head.Next
}

// print list
func pprint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	list1 := newListNode([]int{1, 4, 3, 2, 5, 2})

	fmt.Println(pprint(partition(list1, 3)))
}

func partition(head *ListNode, x int) *ListNode {
	// check
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	head = headDummy
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
