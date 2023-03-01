package main

import "fmt"

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
		// fmt.Println(head.Val)
		head = head.Next
	}
	return res
}

func main() {
	list1 := newListNode([]int{3, 2, 0, -4})
	listWithCycle := buildCycle(list1, 1)
	res := detectCycle(listWithCycle)
	fmt.Println(res.Val)
	fmt.Println(res.Next.Val)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			fast = fast.Next
			slow = head
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	for fast != nil {
		if fast == slow {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}
