package main

import (
	"fmt"
)

// https://leetcode.cn/problems/search-in-rotated-sorted-array/

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。 ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 你可以假设数组中不存在重复的元素。

func main() {
	aa := []int{4, 5, 6, 7, 0, 1, 2}
	bb := []int{1, 3, 5}
	cc := []int{3, 5, 1}
	dd := []int{4, 5, 0, 1, 2}
	fmt.Println(search(aa, 3))
	fmt.Println(search(bb, 1))
	fmt.Println(search(cc, 3))
	fmt.Println(search(dd, 3))
}

func search(nums []int, target int) int {
	// 思路：/ / 两条上升直线，四种情况判断
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		// 相等直接返回
		if nums[mid] == target {
			return mid
		}
		// 将数组从中间分开成左右两部分的时候，一定有一部分的数组是有序的
		// 先判断哪一部分是有序的，然后判断target在不在里面
		if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[end] >= target && nums[mid] <= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}
