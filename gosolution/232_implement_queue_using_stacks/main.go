package main

import "fmt"

// https://leetcode.cn/problems/implement-queue-using-stacks/
// 使用两个栈实现队列
type MyQueue struct {
	inStack     []int
	outputStack []int
}

func main() {
	obj := Constructor()
	x := 1
	obj.Push(x)
	y := 2
	obj.Push(y)
	fmt.Printf("obj %v\n", obj)

	param_2 := obj.Peek()
	fmt.Println("param_2 ", param_2)
	fmt.Printf("obj %v\n", obj)

	param_3 := obj.Pop()
	fmt.Println("param_3 ", param_3)
	fmt.Printf("obj %v\n", obj)

	param_4 := obj.Empty()
	fmt.Println("param_4 ", param_4)
}

func Constructor() MyQueue {
	var queue MyQueue
	return queue
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	var out int
	if len(this.outputStack) == 0 {
		for _, v := range this.inStack {
			this.outputStack = append([]int{v}, this.outputStack...)
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
		out = this.outputStack[len(this.outputStack)-1]
		this.outputStack = this.outputStack[:len(this.outputStack)-1]
	} else {
		out = this.outputStack[len(this.outputStack)-1]
		this.outputStack = this.outputStack[:len(this.outputStack)-1]
	}

	return out
}

func (this *MyQueue) Peek() int {
	var out int
	if len(this.outputStack) != 0 {
		out = this.outputStack[len(this.outputStack)-1]
	} else {
		for _, v := range this.inStack {
			this.outputStack = append([]int{v}, this.outputStack...)
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
		out = this.outputStack[len(this.outputStack)-1]
	}
	return out
}

func (this *MyQueue) Empty() bool {
	if len(this.outputStack) == 0 && len(this.inStack) == 0 {
		return true
	} else {
		return false
	}
}
