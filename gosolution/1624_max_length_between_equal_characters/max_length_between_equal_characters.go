// 给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。
//
// 子字符串 是字符串中的一个连续字符序列。
//
//
//
// 示例 1：
//
// 输入：s = "aa"
// 输出：0
// 解释：最优的子字符串是两个 'a' 之间的空子字符串。
//
// 示例 2：
//
// 输入：s = "abca"
// 输出：2
// 解释：最优的子字符串是 "bc" 。
//
//
// 示例 3：
//
// 输入：s = "cbzxy"
// 输出：-1
// 解释：s 中不存在出现出现两次的字符，所以返回 -1 。
//
//
// 示例 4：
//
// 输入：s = "cabbac"
// 输出：4
// 解释：最优的子字符串是 "abba" ，其他的非最优解包括 "bb" 和 "" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 300
// s 只含小写英文字母
//
//
// Related Topics 哈希表 字符串 👍 17 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func maxLengthBetweenEqualCharacters(s string) int {
	maxLen := -1
	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				maxLen = max(maxLen, j-i-1)
				break
			}
		}
	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// leetcode submit region end(Prohibit modification and deletion)
