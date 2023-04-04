package main

import "fmt"

// https://leetcode.cn/problems/single-number-ii/

func main() {
	aa := []int{2, 2, 1, 2}
	fmt.Println(singleNumber(aa))
}

func singleNumber(nums []int) int {
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
			fmt.Println(sum)
		}
		result ^= (sum % 3) << i
	}
	return result
}

func singleNumberBruteForceSolution(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	return 0
}
