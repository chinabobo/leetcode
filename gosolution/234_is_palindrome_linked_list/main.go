package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
)

// https://leetcode.cn/problems/palindrome-linked-list/
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

func main() {
	list1 := newListNode([]int{1, 2, 2, 3, 3, 1})
	fmt.Println(isPalindrome(list1))
	fmt.Println(isPalindrome2(list1))
}

func isPalindrome(head *ListNode) bool {
	var cou int
	src := head
	for head != nil {
		cou++
		head = head.Next
	}
	if cou&0x1 == 0 { // ou
		return isPalindromeShuang(src, cou)
	} else {
		return isPalindromeOdd(src, cou)
	}
}

func isPalindromeShuang(head *ListNode, cou int) bool {
	tag := false
	mid := cou / 2
	var stack []int
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	for head != nil {
		if len(stack) == mid {
			tag = true
		}
		if len(stack) != 0 && stack[len(stack)-1] == head.Val && tag {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, head.Val)
		}
		fmt.Println(stack)
		head = head.Next
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isPalindromeOdd(head *ListNode, cou int) bool {
	if head.Next == nil {
		return true
	}
	var n int
	m := cou/2 - 1
	fmt.Println(m)

	src := head
	for head != nil {
		if n == m {
			head.Next = head.Next.Next
			fmt.Println(n)
			fmt.Println(head.Val)
			break
		}
		n++
		head = head.Next
	}
	return isPalindromeShuang(src, cou-1)
}

func isPalindrome2(head *ListNode) bool {
	// 1 2 nil
	// 1 2 1 nil
	// 1 2 2 1 nil
	if head == nil {
		return true
	}
	slow := head
	// fast如果初始化为head.Next则中点在slow.Next
	// fast初始化为head,则中点在slow
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	tail := reverse(slow.Next)
	// 断开两个链表(需要用到中点前一个节点)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true

}

func reverse(head *ListNode) *ListNode {
	// 1->2->3
	if head == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}
	return prev
}
