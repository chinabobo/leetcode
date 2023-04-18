package main

// https://leetcode.cn/problems/first-bad-version/

// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

func firstBadVersion(n int) int {
	sta := 1
	end := n
	for sta+1 < end {
		mid := sta + (end-sta)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			sta = mid
		}
	}
	if isBadVersion(sta) {
		return sta
	}
	return end
}

func isBadVersion(n int) bool {
	// fake
	return false
}
