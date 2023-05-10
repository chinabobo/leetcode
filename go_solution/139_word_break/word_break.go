package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))

}
func wordBreak(s string, wordDict []string) bool {
	// f[i] 表示前i个字符是否可以被切分
	// f[i] = f[j] && s[j+1~i] in wordDict
	// f[0] = true
	// return f[len]

	if len(s) == 0 {
		return true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	max, dict := maxLen(wordDict)
	for i := 1; i <= len(s); i++ {
		l := 0
		if i-max > 0 {
			l = i - max
		}
		for j := l; j < i; j++ {
			if f[j] && inDict(s[j:i], dict) {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

func maxLen(wordDict []string) (int, map[string]bool) {
	dict := make(map[string]bool)
	max := 0
	for _, v := range wordDict {
		dict[v] = true
		if len(v) > max {
			max = len(v)
		}
	}
	return max, dict
}

func inDict(s string, dict map[string]bool) bool {
	_, ok := dict[s]
	return ok
}

/*
dp[i] 0,i 是否可以被拆分

if (j == 0 || dp[j-1]) && findWordInDict(s[j:i+1], wordDict)
dp[i] = true
*/

func wordBreakSelf(s string, wordDict []string) bool {
	dp := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(i)
		for j := i; j >= 0; j-- {
			fmt.Println("j : ", j)
			fmt.Println("s[j:i+1] : ", s[j:i+1])
			fmt.Println("findWordInDict(s[j:i+1], wordDict) : ", findWordInDict(s[j:i+1], wordDict))
			if j != 0 {
				fmt.Println("dp[j-1] : ", dp[j-1])
			}
			if (j == 0 || dp[j-1]) && findWordInDict(s[j:i+1], wordDict) {
				fmt.Println("s[j:i+1] : ", s[j:i+1])
				dp[i] = true
				fmt.Println(dp)
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)-1]
}

func findWordInDict(s string, dict []string) bool {
	for _, v := range dict {
		if s == v {
			return true
		}
	}
	return false
}
