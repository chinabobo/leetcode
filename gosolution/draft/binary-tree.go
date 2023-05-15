package main

import (
	"fmt"
	"github.com/chinabobo/leetcode/gosolution/model"
	"github.com/chinabobo/leetcode/gosolution/util/binarytree"
)

func main() {
	root := binarytree.BuildTree([]int{1, 2, 3, 4, 0, 5, 6})
	result := binarytree.ConvertToArr(root)
	fmt.Println(result)
	fmt.Println(BFS(root))
	// fmt.Println(BFS2(root))
}
func BFS(root *model.TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*model.TreeNode, 0)
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
