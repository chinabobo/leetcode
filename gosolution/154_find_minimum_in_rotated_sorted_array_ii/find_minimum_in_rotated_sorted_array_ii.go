package main

import "fmt"

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转 ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 请找出其中最小的元素。(包含重复元素)

func main() {
	aa := []int{2, 2, 2, 0, 1}
	bb := []int{1, 2, 2, 2, 0, 1}
	fmt.Println(findMin(aa))
	fmt.Println(findMin(bb))
}
func findMin(nums []int) int {
	// 思路：跳过重复元素，mid值和end值比较，分为两种情况进行处理
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		// 去除首尾的重复元素
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid := start + (end-start)/2
		// 中间元素和最后一个元素比较（判断中间点落在左边上升区，还是右边上升区）
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
