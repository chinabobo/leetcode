package main

import (
	"fmt"
)

// 给定一个正整数数组 arr，求 arr[i] / arr[j] 的最大值，其中 i < j。
func main() {
	arr := []int{2, 4, 8, 16, 227, 38, 49}
	fmt.Println(findMaxQuotient(arr))
	fmt.Println(findM(arr))
}

func findMaxQuotient(arr []int) int {
	var max int
	if len(arr) < 2 {
		return 0
	}
	for i := range arr {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i]/arr[j] > max {
				max = arr[i] / arr[j]
			}
		}
	}
	return max
}

func findM(src []int) int {
	var res int
	if len(src) < 2 {
		return 0
	}

	for j := 1; j < len(src)-1; j++ {
		for i := 0; i < j; i++ {
			if src[i]/src[j] > res {
				res = src[i] / src[j]
			}
		}
	}
	return res
}
