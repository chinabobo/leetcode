package main

import "fmt"

// https://leetcode.cn/problems/counting-bits/
// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

func main() {
	fmt.Println(countBits(5))
	fmt.Println(countBits2(5))
}

func countBits(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		var item int
		item = count1Bit(i)
		res = append(res, item)
	}
	return res
}

func count1Bit(n int) int {
	var res int
	for n != 0 {
		n &= n - 1
		res++
	}
	return res
}

// Dynamic programming
func countBits2(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 上一个缺1的元素+1即可
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
