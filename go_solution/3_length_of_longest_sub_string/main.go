package main

import "fmt"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func main() {
	str := "uqinntq"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	var window []rune
	var max int
	for _, v := range []rune(s) {
		if NotInWindow(v, window) {
			window = append(window, v)
			fmt.Println(string(window))
		} else {
			if len(window) != 0 {
				window = append(window, v)
				window = window[1:]
			}
			for {
				if !CheckWin(window) {
					window = window[1:]
				} else {
					break
				}
			}

			fmt.Println(string(window))
		}
		if max < len(window) {
			max = len(window)
		}
	}
	return max
}

func NotInWindow(aa int32, win []rune) bool {
	for _, v := range win {
		if aa == v {
			return false
		}
	}
	return true
}

func CheckWin(win []rune) bool {
	for i := range win {
		for j := i + 1; j < len(win); j++ {
			if win[i] == win[j] {
				return false
			}
		}
	}
	return true
}
