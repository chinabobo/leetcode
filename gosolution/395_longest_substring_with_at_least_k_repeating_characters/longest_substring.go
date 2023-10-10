package main

import "strings"

// https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters
// 对于字符串 sss，如果存在某个字符 ch，它的出现次数大于 000 且小于 kkk，
// 则任何包含 ch 的子串都不可能满足要求。也就是说，我们将字符串按照 ch 切分成若干段，
// 则满足要求的最长子串一定出现在某个被切分的段内，而不能跨越一个或多个段。因此，可以考虑分治的方式求解本题。
//

func longestSubstring(s string, k int) (ans int) {
	if s == "" {
		return
	}

	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt[:] {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
