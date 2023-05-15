package main

import "fmt"

/*
https://leetcode.cn/problems/unique-paths-ii

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。
机器人试图达到网格的右下角（在下图中标记为“Finish”）。 问总共有多少条不同的路径？
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

*/

func main() {
	aa := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(aa))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
