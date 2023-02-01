package main

import (
	"fmt"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func main() {
	root := BuildTree([]int{1, 2, 3, 4, 0, 5, 6})
	result := ConvertToArr(root)
	fmt.Println(result)
	fmt.Println(BFS(root))
	// fmt.Println(BFS2(root))
}
func BFS(root *Node) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 为什么要取length？
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}


// BuildTree 输入一个slice ：[3,9,20,0,0,15,7]
func BuildTree(l []int) (root *Node) {
	length := len(l)
	if length == 0 {
		return root
	}

	var nodes = make([]*Node, length)
	root = &Node{
		Val: l[0],
	}
	nodes[0] = root
	//循环输入的数组切片，依次判断每一个节点的左右节点是否存在并创建
	for i := 0; i < length; i++ {
		currentNode := nodes[i]

		if currentNode == nil {
			continue
		}

		leftIndex := 2*i + 1
		if leftIndex < length && l[leftIndex] != 0 {
			currentNode.Left = &Node{
				Val: l[leftIndex],
			}
			nodes[leftIndex] = currentNode.Left
		}

		rightIndex := 2*i + 2
		if rightIndex < length && l[rightIndex] != 0 {
			currentNode.Right = &Node{
				Val: l[rightIndex],
			}
			nodes[rightIndex] = currentNode.Right
		}
	}

	return root
}

// ConvertToArr 树转化成数组
func ConvertToArr(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := indexRecursion(root, 1, 0, []int{})
	// 删除末尾的0
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == 0 {
			result = result[:i]
		} else {
			break
		}
	}

	return result
}

// 索引循环
func indexRecursion(node *Node, level int, index int, result []int) []int {
	//深度为level的树最多有2^level-1个节点，容量不够时扩容依据
	if len(result) < (1<<level - 1) {
		newArr := make([]int, 1<<level-1)
		copy(newArr, result)
		result = newArr
	}
	result[index] = node.Val

	if node.Left != nil {
		result = indexRecursion(node.Left, level+1, 2*index+1, result)
	}

	if node.Right != nil {
		result = indexRecursion(node.Right, level+1, 2*index+2, result)
	}

	return result
}