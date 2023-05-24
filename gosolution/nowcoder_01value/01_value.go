package main

import "fmt"

/*
https://www.nowcoder.com/questionTerminal/16976852ad2f4e26a1ff9f555234cab2

给出一个只包含 0 和 1 的 01 串 s ，下标从 1 开始，设第 i 位的价值为 vali ，则价值定义如下：

1. i=1时:val1 = 1
2. i>1时：
2.1 若 si ≠ si-1 , vali = 1
2.2 若 si = si-1 , vali = vali-1 + 1
字符串的价值等于 val1 + val2 + val3 + ... + valn

你可以删除 s 的任意个字符，问这个串的最大价值是多少。

输入
6
010101
输出
7
说明
删除后的串为0001或0111时有最大价值

输入
20
11111000111011101100
输出
94

输入
4
1100
输出
6
*/
func main() {
	//a := 0
	//b := 0
	//for {
	//	n, _ := fmt.Scan(&a, &b)
	//	if n == 0 {
	//		break
	//	} else {
	//		fmt.Printf("%d\n", a+b)
	//	}
	//}
	//n := 6
	//s := "010101"
	n := 3
	s := "111"
	fmt.Println(calculateMaxValue(n, s))
}

func calculateMaxValue(n int, s string) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		count := 1
		for j := i - 1; j > -1; j-- {
			if s[j] == s[i] {
				if j == 0 && dp[i] < count*(count+1)/2 {
					dp[i] = count * (count + 1) / 2
				}
				count++

			} else {
				if dp[i] < dp[j]+count*(count+1)/2 {
					dp[i] = dp[j] + count*(count+1)/2
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}
