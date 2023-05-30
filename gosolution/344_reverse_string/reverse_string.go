package main

import "fmt"

/*
https://leetcode.cn/problems/reverse-string

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
示例 1：

输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
*/

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	s = []byte("hello")
	reverseString2(s)
	fmt.Println(string(s))

}

/*
recursive
*/
func reverseString(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}
func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}

/*
double pointer
*/
func reverseString2(s []byte) {
	n := len(s)
	for left, right := 0, n-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
