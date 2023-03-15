package main

import "fmt"

func main() {
	str := "aab"
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
				window = window[1:]
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
