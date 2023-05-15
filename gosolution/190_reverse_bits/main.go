package main

import "fmt"

// https://leetcode.cn/problems/reverse-bits/
// 颠倒给定的 32 位无符号整数的二进制位

func main() {
	aa := uint32(43261596)
	fmt.Printf("%b \n", aa)
	fmt.Printf("%b ", reverseBits(aa))
}

func reverseBits(num uint32) uint32 {
	var res uint32
	pom := 31
	for num != 0 {
		res += (num & 1) << pom
		pom--
		num >>= 1
	}
	return res
}
