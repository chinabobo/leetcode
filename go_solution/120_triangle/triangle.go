package main

import (
	"fmt"
	"math"
)

/* https://leetcode.cn/problems/triangle/
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

input:
	triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]

output: 11

explain：
    2
   3 4
  6 5 7
 4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

*/

func main() {
	aa := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotalTraverse(aa))
	fmt.Println(minimumTotalDivide(aa))
	fmt.Println(minimumTotalOpt(aa))
	fmt.Println(minimumTotalBottomUpDp(aa))
	fmt.Println(minimumTotalTopDownDP(aa))
}

// 动态规划和 DFS 区别
// 二叉树 子问题是没有交集，所以大部分二叉树都用递归或者分治法，即 DFS，就可以解决
// 像 triangle 这种是有重复走的情况，子问题是有交集，所以可以用动态规划来解决

// bottom-up DP

func minimumTotalBottomUpDp(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1、状态定义：f[i][j] 表示从i,j出发，到达最后一层的最短路径
	var l = len(triangle)
	var f = make([][]int, l)
	// 2、初始化
	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}
	// 3、递推求解
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
		}
	}
	// 4、答案
	return f[0][0]
}

// top-down DP

// 测试用例：
// [
// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]
// ]
func minimumTotalTopDownDP(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1、状态定义：f[i][j] 表示从0,0出发，到达i,j的最短路径
	var l = len(triangle)
	var f = make([][]int, l)
	// 2、初始化
	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}
	// 递推求解
	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 这里分为两种情况：
			// 1、上一层没有左边值
			// 2、上一层没有右边值
			if j-1 < 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j >= len(f[i-1]) {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	result := f[l-1][0]
	for i := 1; i < len(f[l-1]); i++ {
		result = min(result, f[l-1][i])
	}
	return result
}

// traverse

func minimumTotalTraverse(triangle [][]int) int {
	best := math.MaxInt
	sum := 0
	dfs(0, 0, &sum, &triangle, &best, len(triangle))
	return best
}

// The top-down minimum path

func dfs(x, y int, sum *int, src *[][]int, best *int, n int) {
	if x == n {
		if *sum < *best {
			*best = *sum
		}
		return
	}
	tmp := *sum + (*src)[x][y]
	dfs(x+1, y, &tmp, src, best, n)
	dfs(x+1, y+1, &tmp, src, best, n)
}

// divide & conquer

func minimumTotalDivide(triangle [][]int) int {
	return dfsDiv(0, 0, len(triangle), &triangle)
}

func dfsDiv(x, y, n int, src *[][]int) int {
	if x == n {
		return 0
	}
	return min(dfsDiv(x+1, y, n, src), dfsDiv(x+1, y+1, n, src)) + (*src)[x][y]
}

// divide & conquer optimization

func minimumTotalOpt(triangle [][]int) int {
	hash := make([][]int, len(triangle))
	for i := range hash {
		hash[i] = make([]int, len(triangle[i]))
	}
	return dfsOpt(0, 0, len(triangle), &triangle, &hash)
}

func dfsOpt(x, y, n int, src, hash *[][]int) int {
	if x == n {
		return 0
	}
	if (*hash)[x][y] != 0 {
		return (*hash)[x][y]
	}
	(*hash)[x][y] = min(dfsOpt(x+1, y, n, src, hash), dfsOpt(x+1, y+1, n, src, hash)) + (*src)[x][y]
	return (*hash)[x][y]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
