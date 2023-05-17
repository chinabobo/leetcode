package main

import "strings"

/*
https://leetcode.cn/problems/check-if-binary-string-has-at-most-one-segment-of-ones
给你一个二进制字符串 s ，该字符串 不含前导零 。

如果 s 包含 零个或一个由连续的 '1' 组成的字段 ，返回 true 。否则，返回 false 。


示例 1：

输入：s = "1001"
输出：false
解释：由连续若干个 '1' 组成的字段数量为 2，返回 false
示例 2：

输入：s = "110"
输出：true

*/
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
