package main

import "fmt"

func main() {
	aa := []int{2, 4, 7, 4, 1, 3}
	fmt.Println(MergeSort(aa))
}

func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：divide 分为两段
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并两段数据
	result := merge(left, right)
	return result
}

func merge(left, right []int) (result []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 注意不能越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
