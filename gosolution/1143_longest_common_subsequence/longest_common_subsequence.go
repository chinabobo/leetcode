package main

import "fmt"

/*
https://leetcode.cn/problems/longest-common-subsequence/

给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
*/
func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

/*
相等取左上元素+1，否则取左或上的较大值
从 1 开始遍历到最大长度
索引需要减一

		a	b	c	d	e
	a	0	0	0	0	0	0
	c	0	1	1	1	1	1
   	e	0	1	1	2	2	2
		0	1	1	2	2	3

*/
func longestCommonSubsequence(a string, b string) int {
	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[len(a)][len(b)]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
