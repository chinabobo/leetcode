package binarytree

import . "github.com/chinabobo/leetcode/model"

func BuildTree(l []int) (root *TreeNode) {
	length := len(l)
	if length == 0 {
		return root
	}

	var nodes = make([]*TreeNode, length)
	root = &TreeNode{
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
			currentNode.Left = &TreeNode{
				Val: l[leftIndex],
			}
			nodes[leftIndex] = currentNode.Left
		}

		rightIndex := 2*i + 2
		if rightIndex < length && l[rightIndex] != 0 {
			currentNode.Right = &TreeNode{
				Val: l[rightIndex],
			}
			nodes[rightIndex] = currentNode.Right
		}
	}

	return root
}

// ConvertToArr 树转化成数组
func ConvertToArr(root *TreeNode) []int {
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
func indexRecursion(node *TreeNode, level int, index int, result []int) []int {
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
