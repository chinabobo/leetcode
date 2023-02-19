package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := CreateNodeInsterOnHead([]int{5, 3}) // 3, 5

	reverseBetween(head, 1, 2)
}

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

// https://leetcode.cn/problems/reverse-linked-list-ii/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
