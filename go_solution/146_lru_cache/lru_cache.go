package main

import "fmt"

/*
https://leetcode.cn/problems/lru-cache/
*/
func main() {
	obj := Constructor(2)

	obj.Get(2)
	obj.Put(2, 6)
	obj.Get(1)
	obj.Put(1, 5)
	obj.Put(1, 2)
	obj.Get(1)
	fmt.Println(obj.Get(2))
}

type Node struct {
	key, value int
	pre, next  *Node
}

type LRUCache struct {
	n        int
	capacity int
	m        map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	res.head.next = res.tail
	res.tail.pre = res.head
	return res
}

func initNode(k, v int) *Node {
	return &Node{
		key:   k,
		value: v,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	node := this.m[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok {
		node := initNode(key, value)
		this.m[key] = node
		this.addToHead(node)
		this.n++
		if this.n > this.capacity {
			removed := this.removeNode(this.tail.pre)
			delete(this.m, removed.key)
			this.n--
		}
	} else {
		node := this.m[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) *Node {
	node.pre.next = node.next
	node.next.pre = node.pre
	return node
}

func (this *LRUCache) addToHead(node *Node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}
