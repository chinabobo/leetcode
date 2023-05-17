package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
)

// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 头插法
func CreateNodeInsterOnHead(listData []int) *ListNode {
	if len(listData) == 0 {
		return &ListNode{Val: 0}
	}
	head := &ListNode{Val: listData[0]}
	for _, v := range listData[1:] {
		node := &ListNode{Val: v}
		node.Next = head
		head = node
	}
	return head
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

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	count := 0
	newHead := head
	for head != nil {
		count++
		head = head.Next
	}

	res := make([]int, count)
	i := 0
	for newHead != nil {
		res[count-i-1] = newHead.Val
		i++
		newHead = newHead.Next
	}

	return res
}

func main() {
	list1 := CreateNodeInsterOnHead([]int{4, 2, 1}) // 1, 2, 4
	list2 := CreateNodeInsterOnHead([]int{4, 3, 1}) // 1, 3, 4

	fmt.Println(pprint(mergeTwoListsMyAns(list1, list2)))
	// fmt.Println(reversePrint(mergeTwoLists(list1, list2)))
}

func mergeTwoListsMyAns(list1 *ListNode, list2 *ListNode) *ListNode {
	p := list1
	q := list2
	var newl *ListNode
	nood := ListNode{Val: 0}
	newl = &nood
	pre := newl
	if p == nil {
		return q
	}
	if q == nil {
		return p
	}
	for {
		if q == nil && p == nil {
			break
		}
		if p == nil && q != nil {
			pre.Next = q
			q = q.Next
			pre = pre.Next
			continue
		}
		if q == nil && p != nil {
			pre.Next = p
			p = p.Next
			pre = pre.Next
			continue
		}
		if p.Val < q.Val {
			pre.Val = p.Val
			pre.Next = p
			if p.Next == nil {
				pre.Next = q
				break
			}
			p = p.Next
			pre = pre.Next
		} else {
			pre.Val = q.Val
			pre.Next = q
			if q.Next == nil {
				pre.Next = p
				break
			}
			q = q.Next
			pre = pre.Next
		}

	}
	return newl
}

// 思路：通过 dummy node 链表，连接各个元素
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	// 连接l1 未处理完节点
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	// 连接l2 未处理完节点
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}
