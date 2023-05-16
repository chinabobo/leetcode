package binarytree

import (
	. "github.com/chinabobo/leetcode/model"
)

func preorderTraversalRecursive(root *TreeNode, res *[]int) []int {
	if root == nil {
		return *res
	}
	*res = append(*res, root.Val)
	preorderTraversalRecursive(root.Left, res)
	preorderTraversalRecursive(root.Right, res)
	return *res
}

func inorderTraversalRecursive(root *TreeNode, res *[]int) []int {
	if root == nil {
		return *res
	}
	inorderTraversalRecursive(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalRecursive(root.Right, res)
	return *res
}

func postorderTraversalRecursive(root *TreeNode, res *[]int) []int {
	if root == nil {
		return *res
	}
	postorderTraversalRecursive(root.Left, res)
	postorderTraversalRecursive(root.Right, res)
	*res = append(*res, root.Val)
	return *res
}
