package main

/*
https://leetcode.cn/problems/subarray-sum-equals-k/solution/he-wei-kde-zi-shu-zu-by-leetcode-solution
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

input : [1, 1, 1] 2
output : 2


*/

func subarraySum(nums []int, k int) int {
	// Time complexity O(n)
	// Space complexity O(n)

	// pre[i] 为 0 - i 所有数之和
	// pre[i] = pre[i-1] + nums[i]
	// j - i 这个子数组的和为k => pre[i] - pre[j-1] == k => pre[j-1] == pre[i] - k
	// 所以以 i 结尾的和为 k 的连续子数组个数时只要统计有多少个前缀和为 pre[i] - k 的 pre[j] 即可

	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func subarraySumBruteForceSolution(nums []int, k int) int {
	// Time complexity O(n2)
	// Space complexity O(1)
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySumSelf(nums []int, k int) int {
	var cnt int
	for i := 0; i < len(nums); i++ {
		tt := 0
		for j := 0; i+j < len(nums); j++ {
			tt += nums[i+j]
			if tt == k {
				cnt++
			}
		}

	}
	return cnt
}
