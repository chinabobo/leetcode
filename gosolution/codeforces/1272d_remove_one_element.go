package main

/*
https://codeforces.com/problemset/problem/1272/D

题意：第一行给出一个数n，接下来给出n个数的序列。求在该序列中删去任意一个数后所得序列中的最长连续的递增子序列长度

input
5
1 2 5 3 4
output
4


ans 1 : https://blog.csdn.net/qq_41280600/article/details/103519859
ans 2 : https://blog.csdn.net/Mr_Kingk/article/details/103606895

dp【i】【0】: 代表以a[i]结尾的字串没有删除任何元素的最大长度。
dp【i】【1】:代表以a[i]结尾的字串删除了一个元素的最大长度。
若当前的a【i】 > a【i-1】那么可以更新2种状态
(1) dp[i][0] = max(dp[i][0], dp[i - 1][0] + 1); 代表当前没有选择删除元素，那么只能由上一个没有选择删除元素的状态转移。
(2) dp[i][1] = max(dp[i][1], dp[i - 1][1] + 1); , 代表当前选择了删除元素，那么由上一状态已经删除了的转移过来。
若a【i】 > a【i-2】 那么可以选择是否删除a【i-1】这个元素
dp[i][1] = max(dp[i][1], dp[i - 2][0] + 1); 这个状态就由dp【i-2】【0】 之前没有删除过元素的状态转移过来。

*/

func findMaxLen(src []int) int {
	var ans int
	dp := make([][2]int, len(src))
	dp[1][0] = 1
	for i := 2; i < len(src); i++ {
		dp[i][0] = 1
		dp[i][1] = 1
		if src[i] > src[i-1] {
			dp[i][0] = max(dp[i][0], dp[i-1][0]+1)
			dp[i][1] = max(dp[i][1], dp[i-1][1]+1)
		}
		if src[i] > src[i-2] {
			dp[i][1] = max(dp[i][1], dp[i-2][0]+1)
		}
		ans = max(ans, max(dp[i][0], dp[i][1]))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
