package binarytree

import (
	"fmt"
	. "github.com/chinabobo/leetcode/gosolution/model"
)

func preorderTraversalRecursive(root *TreeNode) {
	for root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversalRecursive(root.Left)
	preorderTraversalRecursive(root.Right)
}

func inorderTraversalRecursive(root *TreeNode) {
	for root == nil {
		return
	}
	inorderTraversalRecursive(root.Left)
	fmt.Println(root.Val)
	inorderTraversalRecursive(root.Right)
}

func postorderTraversalRecursive(root *TreeNode) {
	for root == nil {
		return
	}
	postorderTraversalRecursive(root.Left)
	postorderTraversalRecursive(root.Right)
	fmt.Println(root.Val)
}
