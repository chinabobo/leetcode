package main

import "fmt"

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 请找出其中最小的元素。
func main() {
	// aa := []int{3, 4, 5, 1, 2}
	// bb := []int{5, 1, 2, 3, 4}
	cc := []int{2, 3, 4, 5, 1}
	// fmt.Println(findMin(aa))
	// fmt.Println(findMin(bb))
	fmt.Println(findMin(cc))
}

func findMin(nums []int) int {
	// 思路：最后一个值作为target，然后往左移动，最后比较start、end的值
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		// 最后一个元素值为target
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

// 初次的想法，判断前后，仔细思考后发现是不行的
func findMinSelfFailed(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		} else if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
			end = mid
		} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
