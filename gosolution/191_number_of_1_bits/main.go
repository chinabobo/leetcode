package main

import (
	"fmt"
)

// https://leetcode.cn/problems/number-of-1-bits/
// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

func main() {
	aa := uint32(11)
	bb := uint32(4294967293)
	fmt.Println(hammingWeight(aa))
	fmt.Println(hammingWeight(bb))
}

func hammingWeight(num uint32) int {
	var count int
	x := num
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
