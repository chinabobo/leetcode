package main

import (
	"fmt"
	"github.com/chinabobo/leetcode/util"
)

func main() {
	aa := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(aa))
}

/*
https://leetcode.cn/problems/container-with-most-water

给定一个长度为 n 的整数数组 height 。有n条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

思路：
双指针，由于可容纳水的高度由两板中的 短板 决定

始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。
在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽 底边宽度 变短
如果移动短板， min(h(i), h(j)) 有可能增加
如果移动长板， 水槽面积一定减少
所以每次移动都是短板
*/
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var res int
	for i < j {
		if res < util.Min(height[i], height[j])*(j-i) {
			res = util.Min(height[i], height[j]) * (j - i)
		}
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return res
}
