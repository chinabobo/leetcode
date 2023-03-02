package main

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	minVal := ms.GetMin()
	if val < minVal {
		ms.min = append(ms.min, val)
	} else {
		ms.min = append(ms.min, minVal)
	}
	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.min = ms.min[:len(ms.min)-1]
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.min) == 0 {
		return 1 << 31
	}
	return ms.min[len(ms.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
