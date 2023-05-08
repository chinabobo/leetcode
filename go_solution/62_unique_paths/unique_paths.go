package main

import "fmt"

/*
https://leetcode.cn/problems/unique-paths/

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。
机器人试图达到网格的右下角（在下图中标记为“Finish”）。 问总共有多少条不同的路径？

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下

f(0, 0) = 1

f(i, j) = f(i-1, j) + f(i, j-1)

*/

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
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
