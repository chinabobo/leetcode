package main

import (
	"fmt"
)

// https://leetcode.cn/problems/sort-list/
// 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
// 思路：归并排序，找中点和合并操作

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

	fmt.Println(pprint(sortList(list1)))
}

func sortList(head *ListNode) *ListNode {
	// 思路：归并排序，找中点和合并操作
	return mergeSort(head)
}

func findMiddle(head *ListNode) *ListNode {
	// 1->2->3->4->5
	slow := head
	fast := head.Next
	// 快指针先为nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
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
func mergeSort(head *ListNode) *ListNode {
	// 如果只有一个节点 直接就返回这个节点
	if head == nil || head.Next == nil {
		return head
	}
	// find middle
	middle := findMiddle(head)
	fmt.Println(middle.Val)
	// 断开中间节点
	tail := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	result := mergeTwoLists(left, right)
	return result
}
