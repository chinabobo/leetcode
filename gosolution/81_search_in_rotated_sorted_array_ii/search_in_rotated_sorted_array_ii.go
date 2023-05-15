package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。 ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。(包含重复元素)

func main() {
	aa := []int{1, 3, 5}
	fmt.Println(search(aa, 1))
}

func search(nums []int, target int) bool {
	// 思路：/ / 两条上升直线，四种情况判断，并且处理重复数字
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		// 处理重复数字
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		mid := start + (end-start)/2
		// 相等直接返回
		if nums[mid] == target {
			return true
		}
		// 判断在那个区间，可能分为四种情况
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
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}

func searchSelf(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for nums[start] == nums[start+1] {
			start = start + 1
			if start+1 >= end {
				break
			}
		}
		for nums[end] == nums[end-1] {
			end = end - 1
			if start+1 >= end {
				break
			}
		}
		if start+1 >= end {
			break
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[start] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[mid] < nums[end] {
			if nums[mid] < target && nums[end] >= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	} else {
		return false
	}
}
