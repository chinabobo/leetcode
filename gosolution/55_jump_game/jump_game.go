package main

import "fmt"

/*
https://leetcode.cn/problems/jump-game/

给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
*/

func main() {
	// aa := []int{2, 3, 1, 1, 4}
	bb := []int{3, 2, 1, 0, 4}
	// fmt.Println(canJump2(aa))
	fmt.Println(canJump2(bb))
}

/*
思路：
每次跳的落点都是起跳点true
如果一个位置能够到达，那么这个位置左侧所有位置都能到达
*/

func canJump2(nums []int) bool {
	var farthest int
	for i := range nums {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
思路：看最后一跳
状态：dp[i] 表示是否能从0跳到i
dp[i] = OR(dp[j], j < i && j 能跳到 i) 判断之前所有的点最后一跳是否能跳到当前点
dp[0] = 0
dp[n-1]
*/
func canJump0(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	dp := make([]bool, n)

	dp[0] = true

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && nums[j]+j >= i {
				dp[i] = true
			}
		}
	}
	return dp[n-1]
}
