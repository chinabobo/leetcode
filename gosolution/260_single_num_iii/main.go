package main

import "fmt"

// https://leetcode.cn/problems/single-number-iii/

// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

func main() {
	aa := []int{1, 2, 1, 3, 2, 5} // output {3, 5}
	fmt.Println(singleNumber2(aa))
}

func singleNumber(nums []int) []int {
	// a=a^b
	// b=a^b
	// a=a^b
	// 关键点怎么把a^b分成两部分,方案：可以通过diff最后一个1区分

	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

func singleNumber2(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum //  只保留 x 的二进制表示中最后一个 1
	// lsb := (xorSum & (xorSum - 1)) ^ xorSum  // 这两种计算lowbit方式是一样的，都可以计算出，而且结果一样
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
