package main

/**
在 n 个物品中挑选若干物品装入背包，最多能装多满？
假设背包的大小为 m，每个物品的大小为Ai（每个物品只能选择一次且物品大小均为正整数）

样例 1：
输入：
a = [3,4,8,5]
backpack size = 10
输出： 9
解释： 装4和5.

样例 2：
输入：
a = [2,3,5,7]
backpack size = 12
输出： 12 装5和7
 * @param m: An integer m denotes the size of a backpack
 * @param a: Given n items with size A[i]
 * @return: The maximum size
*/
func backpack(m int, a []int) int {
	// f[i][j] 前i个物品，是否能装j
	// f[i][j] = f[i-1][j]
	// if j-a[i-1] >= 0 && f[i-1][j-a[i-1]]
	// f[i][j] = true

	// f[0][0]=true f[...][0]=true
	// f[n][X]
	dp := make([][]bool, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(a); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-a[i-1] >= 0 && dp[i-1][j-a[i-1]] {
				dp[i][j] = true
			}
		}
	}
	for i := m; i >= 0; i-- {
		if dp[len(a)][i] {
			return i
		}
	}
	return 0
}
