package main

import "fmt"

//  https://www.lintcode.com/problem/search-for-a-range/description
//  给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
//  如果目标值不在数组中，则返回[-1, -1]

// 数组 = [5, 7, 7, 8, 8, 10]
// target = 8

// output [3,4]

// 思路：核心点就是找第一个 target 的索引，和最后一个 target 的索引，所以用两次二分搜索分别找第一次和最后一次的位置

func main() {
	aa := []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	fmt.Println(searchRange(aa, 5))
}

/**
 * @param a: an integer sorted array
 * @param target: an integer to be inserted
 * @return: a list of length 2, [index1, index2]
 */
func searchRange(a []int, target int) []int {
	if len(a) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	start := 0
	end := len(a) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] > target {
			end = mid
		} else if a[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向左找，就能找到第一个目标值的位置
			end = mid
		}
	}
	// 搜索左边的索引
	if a[start] == target {
		result[0] = start
	} else if a[end] == target {
		result[0] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	start = 0
	end = len(a) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] > target {
			end = mid
		} else if a[mid] < target {
			start = mid
		} else {
			// 如果相等，应该继续向右找，就能找到最后一个目标值的位置
			start = mid
		}
	}
	// 搜索右边的索引
	if a[end] == target {
		result[1] = end
	} else if a[start] == target {
		result[1] = start
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}
