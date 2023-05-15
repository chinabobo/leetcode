package main

import (
	. "../model"
)

/*
https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
给定一棵二叉搜索树，请找出其中第 k 大的节点的值。

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4

dfs 解法

*/

func kthLargest(root *TreeNode, k int) int {
	count := 0
	ans := 0
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Right)

		count++
		if count == k {
			ans = n.Val
			return
		}

		dfs(n.Left)
	}
	dfs(root)
	return ans
}
