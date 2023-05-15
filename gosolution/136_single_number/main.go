package main

import "fmt"

// https://leetcode.cn/problems/single-number/

func main() {
	aa := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(aa))
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

func singleNumberBruteForceSolution(nums []int) int {
	tmap := make(map[int]bool, 0)
	for _, v := range nums {
		if _, ok := tmap[v]; ok {
			tmap[v] = false
		} else {
			tmap[v] = true
		}
	}

	for i := range tmap {
		if tmap[i] {
			return i
		}
	}
	return 0
}
