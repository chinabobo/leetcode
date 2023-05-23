package main

import (
	"github.com/chinabobo/leetcode/util"
)

/*
https://leetcode.cn/problems/edit-distance/

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数
你可以对一个单词进行如下三种操作： 插入一个字符 删除一个字符 替换一个字符

思路
相等则不需要操作，否则取删除、插入、替换最小操作次数的值+1

dp[i][j] 表示a字符串的前i个字符编辑为b字符串的前j个字符最少需要多少次操作
相等则不需要操作 dp[i][j] = dp[i-1][j-1]
否则取删除、插入、替换最小操作次数的值+1 dp[i][j] = util.Min(util.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1

		h	o	r	s	e
	r	0	1	2	3	4	5
	o	1	1	2	2	3	4
	s	2	2	1	2	3	4
		3	3	2	2	3	4
*/
func minDistance(word1 string, word2 string) int {
	//
	// dp[i][j] = OR(dp[i-1][j-1]，a[i]==b[j],min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1)
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 相等则不需要操作
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { // 否则取删除、插入、替换最小操作次数的值+1
				dp[i][j] = util.Min(util.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}
