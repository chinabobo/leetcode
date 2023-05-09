package main

import "fmt"

/*
https://leetcode.cn/problems/palindrome-partitioning-ii/

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。 返回符合要求的最少分割次数。
输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
*/

func main() {

	// fmt.Println(minCut("a")) // 0
	fmt.Println(minCut("ab"))     // 1
	fmt.Println(minCut("aa"))     // 0
	fmt.Println(minCut("aab"))    // 1
	fmt.Println(minCut("cdd"))    // 1
	fmt.Println(minCut("efe"))    // 0
	fmt.Println(minCut("abccbc")) // 2
	fmt.Println(minCut("abaaba")) // 0
	fmt.Println(minCut("cbbbcc")) // 1
}

/*
f(i) 从起点到 i 最少分割次数

f(i) = min(f(j) + 1)
其中 j < i && s[j+1, i] is palindrome

f(0) = 0
*/
func minCut(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if isPalindrome(s, j, i) {
				if j == 0 {
					dp[i] = 0
				}
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPalindrome(src string, j, i int) bool {
	for j < i {
		if src[j] != src[i] {
			return false
		}
		j++
		i--
	}
	return true
}
