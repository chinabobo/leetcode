package main

import "fmt"

/*
https://leetcode.cn/problems/minimum-path-sum/

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

f(i, j) 左上到右下和的最小值

f(0, 0) = 0
f(0, 1) = grid(0, 0) + grid(0, 1)
f(1, 0) = grid(0, 0) + grid(1, 0)

f(i, j) = min(f(i-1, j), f(i, j-1), f(i-1, j-1)) + grid(i, j)


*/

func main() {
	aa := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	bb := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(minPathSum(aa))
	fmt.Println(minPathSum(bb))
}

func minPathSum(grid [][]int) int {
	// f[i][j] 表示i,j到0,0的和最小
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 复用原来的矩阵列表 省下空间
	// 初始化：f[i][0]、f[0][j]
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func minPathSumSelf(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// dp[0][0] = grid[0][0]
	// dp[0][1] = grid[0][0] + grid[0][1]
	// dp[1][0] = grid[0][0] + grid[1][0]

	for i := 0; i <= n-1; i++ {
		for j := 0; j <= m-1; j++ {
			dp[i][j] = grid[i][j]
			if i == 0 && j == 0 {
				dp[i][j] += 0
				continue
			}
			if i == 0 {
				dp[i][j] += dp[i][j-1]
				continue
			}
			if j == 0 {
				dp[i][j] += dp[i-1][j]
				continue
			}
			dp[i][j] += min(dp[i-1][j], dp[i][j-1])
		}
	}
	fmt.Printf("%v \n", dp)
	return dp[n-1][m-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
