package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
)

// https://leetcode.cn/problems/linked-list-cycle/
// 给定一个链表，判断链表中是否有环。
// 快慢指针，指针相遇就有

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

func buildCycle(head *ListNode, pos int) *ListNode {
	var count int
	var tmp *ListNode
	count = 0
	src := head
	for head != nil {
		if count == pos {
			tmp = head
		}
		if head.Next == nil {
			head.Next = tmp
			break
		}
		head = head.Next
		count++
	}
	return src
}

// print list
func pprint(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		fmt.Println(head.Val)
		head = head.Next
	}
	return res
}

func main() {
	list1 := newListNode([]int{3, 2, 0, -4})
	fmt.Println(hasCycle(list1))
	list := buildCycle(list1, 1)
	fmt.Println(hasCycle(list))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
