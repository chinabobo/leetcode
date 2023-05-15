package main

import "fmt"

/*

https://leetcode.cn/problems/longest-increasing-subsequence/
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

input: nums = [10,9,2,5,3,7,101,18]
output: 4


explain: 最长递增子序列是 [2,3,7,101]，因此长度为 4

*/

func main() {
	// aa := []int{4, 10, 4, 3, 8, 9} // 3
	aa := []int{10, 9, 2, 5, 3, 7, 101, 18} // 4
	// fmt.Println(lengthOfLISDP(aa))
	fmt.Println(lengthOfLISGreedy(aa))
}

/*
greedy + binary search
考虑一个简单的贪心，如果我们要使上升子序列尽可能的长，则我们需要让序列上升得尽可能慢，因此我们希望每次在上升子序列最后加上的那个数尽可能的小。
维护一个数组d[i],表示长度为 i 的最长上升子序列的末尾元素的最小值
注意 d[i] 是单调递增的。
依次遍历数组nums的每个元素，并更新数组 d

		 nums[i] > ll[len(ll)-1] : nums[i] 直接加入d的末尾
		/
nums[i]
		\
		 在d中二分查找，找到第一个比nums[i] 小的 d[k], 更新 d[k+1] = nums[i], 如果 nums[i] 比数组 d 中所有的元素都小的话，更新 d[0] =  nums[i]

*/

func lengthOfLISGreedy(nums []int) int {
	ll := make([]int, 0)
	ll = append(ll, nums[0])
	for i := range nums {
		fmt.Println(ll)
		fmt.Println(nums[i])
		if nums[i] > ll[len(ll)-1] {
			ll = append(ll, nums[i])
			continue
		}
		check(ll, nums[i])
	}
	return len(ll)
}

func check(src []int, a int) {
	n := len(src)
	start := 0
	end := n - 1
	for start < end-1 {
		mid := start + (end-start)/2
		if src[mid] < a {
			start = mid
		} else if src[mid] > a {
			end = mid
		} else if src[mid] == a {
			return
		}
	}
	if src[start] == a || src[end] == a {
		return
	}
	if src[end] < a {
		return
	} else if src[start] > a {
		src[start] = a
	} else {
		src[end] = a
	}
	return
}

/*

nums : 0, 3, 1, 6, 2, 2, 7
dp :   1, 2, 2, 3, 3, 3, 4


f(i) i位置的最长递增子序列
f(0) = 1

f[i] 表示从0开始到i结尾的最长序列长度
f[i] = max(f[j])+1 ,a[j]<a[i]

*/

func lengthOfLISDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxx := 0
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < n; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxx {
			maxx = dp[i]
		}
	}
	fmt.Println(dp)
	return maxx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
