package main

import "github.com/chinabobo/leetcode/util"

/*
有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值. 问最多能装入背包的总价值是多大?

样例 1：

输入：
m = 10
A = [2, 3, 5, 7]
V = [1, 5, 2, 4]
输出： 9
解释： 装入 A[1] 和 A[3] 可以得到最大价值, V[1] + V[3] = 9

样例 2：
输入：
m = 10
A = [2, 3, 8]
V = [2, 5, 8]
输出： 10
解释： 装入 A[0] 和 A[2] 可以得到最大价值, V[0] + V[2] = 10
*/

func backpackII(m int, a []int, v []int) int {
	// f[i][j] 前i个物品，装入j背包 最大价值
	// f[i][j] =max(f[i-1][j] ,f[i-1][j-A[i]]+V[i]) 是否加入A[i]物品
	// f[0][0]=0 f[0][...]=0 f[...][0]=0
	dp := make([][]int, len(a)+1)

	for i := 0; i < len(a)+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-a[i-1] >= 0 {
				dp[i][j] = util.Max(dp[i-1][j], dp[i-1][j-a[i-1]]+v[i-1])
			}
		}
	}
	return dp[len(a)][m]

}
