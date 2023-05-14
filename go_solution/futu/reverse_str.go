package main

import "fmt"

func main() {
	fmt.Println(reverseStr("hello"))
}

func reverseStr(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseStr(s[1:]) + string(s[0])
}
