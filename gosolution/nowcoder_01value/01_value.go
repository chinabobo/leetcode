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
	// a := 0
	// b := 0
	// for {
	//	n, _ := fmt.Scan(&a, &b)
	//	if n == 0 {
	//		break
	//	} else {
	//		fmt.Printf("%d\n", a+b)
	//	}
	// }
	// n := 6
	// s := "010101"
	n := 20
	s := "11111111111111111111"
	fmt.Println(calculateMaxValue(n, s))
}

/*
从第二项开始遍历字符串。
对每一项，其价值最小为上一项价值+1，如221。
所以先将其价值预设为最小值，val[i]=val[i-1]+1。并设置一个计数，cnt=1。
从i-1项开始，从后向前遍历字符串。当遇到与i项相同的字符时，令cnt++。
当遇到不同字符时，最大价值分两种情况，
一种是最大价值为前半段采用该不同字符的价值，而后半段则由之后与i项相同的连续字符串组成，即val[j] + cnt * (cnt + 1) / 2
其中cnt * (cnt + 1) / 2为连续cnt个字符的价值，是基本的等差数列求和公式。
另一种情况是当前最大价值字符串中不需要该不同的字符，将其从字符串中剔除，则其不对字符串价值产生影响。val[i]不变。
当全部遍历后，计算将前面所有不同字符剔除的价值，判断其与当前价值，即保留不同字符组合而成的价值的大小关系，取最大值作为最大价值。
*/
func calculateMaxValue(n int, s string) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + 1
		count := 1
		for j := i - 1; j > -1; j-- {
			if s[j] == s[i] {
				// if j == 0 && dp[i] < count*(count+1)/2 {
				// 	dp[i] = count * (count + 1) / 2
				// }
				count++

			} else {
				if dp[i] < dp[j]+count*(count+1)/2 {
					dp[i] = dp[j] + count*(count+1)/2
				}
			}
		}
		if dp[i] < count*(count+1)/2 {
			dp[i] = count * (count + 1) / 2
		}

	}

	fmt.Println(dp)
	return dp[n-1]
}
