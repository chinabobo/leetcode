package binarytree

import (
	. "../../model"
)

func preorderTraversalTopdown(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

// V1：深度遍历，结果指针作为参数传入到函数内部
func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

// V2：通过分治法遍历
func preorderTraversalBottomup(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}
func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	// 返回条件(null & leaf)
	if root == nil {
		return result
	}
	// 分治(Divide)
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并结果(Conquer)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// BFS 层次遍历
func levelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
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
