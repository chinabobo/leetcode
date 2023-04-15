package main

import "fmt"

func main() {
	aa := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search1(aa, 9))
	fmt.Println(search2(aa, 9))
}

func search2(nums []int, target int) int {
	// 1、初始化start、end
	start := 0
	end := len(nums) - 1
	// 2、处理for循环
	for start+1 < end {
		mid := start + (end-start)/2
		// 3、比较a[mid]和target值
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	// 4、最后剩下两个元素，手动判断
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

func search1(nums []int, target int) int {
	return searchMid(0, len(nums)-1, target, nums)
}

func searchMid(s, e, target int, nums []int) int {
	mid := (s + e) / 2
	if nums[mid] < target {
		return searchMid(mid, e, target, nums)
	} else if nums[mid] < target {
		return searchMid(s, mid, target, nums)
	} else if nums[mid] == target {
		return mid
	}
	return -1
}
