package main

import "fmt"

/*
https://leetcode.cn/problems/jump-game-ii/
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。

*/
func main() {
	aa := []int{
		2, 3, 1, 1, 4,
	}
	fmt.Println(jump2(aa)) // 2
	fmt.Println(jump(aa))  // 2
}

// v2 动态规划+贪心优化
func jump2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		// 取第一个能跳到当前位置的点即可
		// 因为跳跃次数的结果集是单调递增的，所以贪心思路是正确的
		idx := 0
		for idx < n && idx+nums[idx] < i {
			idx++
		}
		dp[i] = dp[idx] + 1
	}
	fmt.Println(dp)
	return dp[n-1]
}

/*
dp[i] = dp[j], nums[j] + j >= i,min(dp[j]+1)
*/
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
