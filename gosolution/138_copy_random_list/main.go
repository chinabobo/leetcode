package main

// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。 要求返回这个链表的 深拷贝。
// 思路：1、hash 表存储指针，2、复制节点跟在原节点后面

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 复制节点，紧挨到到后面
	// 1->2->3  ==>  1->1'->2->2'->3->3'
	cur := head
	for cur != nil {
		clone := &Node{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	// 处理random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离两个链表
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	// 原始链表头：head 1->2->3
	// 克隆的链表头：cloneHead 1'->2'->3'
	return cloneHead
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cachedNode[node]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList2(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}
