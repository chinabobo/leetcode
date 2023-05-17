package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
)

// https://leetcode.cn/problems/reorder-list/
// 给定一个单链表 L：L→L→…→L__n→L 将其重新排列后变为： L→L__n→L→L__n→L→L__n→…
// 思路：找到中点断开，翻转后面部分，然后合并前后两个链表
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
	list1 := newListNode([]int{1, 2, 3, 4, 5})
	reorderList(list1)
	fmt.Println(pprint(list1))
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMiddle(head)
	tail := reserve(mid.Next)
	mid.Next = nil
	head = merge(head, tail)
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reserve(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	toggle := true
	for l1 != nil && l2 != nil {
		if toggle {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		toggle = !toggle
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}
