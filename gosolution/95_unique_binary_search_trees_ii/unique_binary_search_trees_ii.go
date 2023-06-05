package main

import (
	"fmt"
	. "github.com/chinabobo/leetcode/model"
	"github.com/chinabobo/leetcode/util/binarytree"
)

/**
https://leetcode.cn/problems/unique-binary-search-trees-ii

给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

*/
func main() {
	for _, v := range generateTrees(3) {
		result := binarytree.ConvertToArr(v)
		fmt.Println(result)
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)

}
func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	ans := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		// 递归生成所有左右子树
		lefts := generate(start, i-1)
		rights := generate(i+1, end)
		// 拼接左右子树后返回
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{Val: i}
				root.Left = lefts[j]
				root.Right = rights[k]
				ans = append(ans, root)
			}
		}
	}
	return ans
}
