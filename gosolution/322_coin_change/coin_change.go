package main

import (
	"github.com/chinabobo/leetcode/util"
	"math"
)

/*
https://leetcode.cn/problems/coin-change

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

*/
func coinChange(coins []int, amount int) int {
	// 状态 dp[i]表示金额为i时，组成的最小硬币个数
	// 推导 dp[i]  = min(dp[i-1], dp[i-2], dp[i-5])+1, 前提 i-coins[j] > 0
	// 初始化为最大值 dp[i]=amount+1
	// 返回值 dp[n] or dp[n]>amount =>-1
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = util.Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]

}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for j := range coins {
			if i-coins[j] < 0 || 1+dp[i-coins[j]] == math.MinInt64 {
				continue
			}

			dp[i] = util.Min(dp[i], 1+dp[i-coins[j]])
		}

	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}

	return dp[amount]
}
