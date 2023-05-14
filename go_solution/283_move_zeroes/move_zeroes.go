package main

import "fmt"

/*
https://leetcode.cn/problems/move-zeroes/

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/
func main() {
	aa := []int{
		0, 1, 0, 3, 12,
	}
	moveZeroes(aa)
	fmt.Println(aa)
}

/*
i和j的相对位置是：【已经处理好的数据】 j 【全是需要被交换的0】 i 【未处理的数据】
在每一次循环前，j 的左边全部都是不等于0的

起始j为0，明显满足
此后每一次循环中，若nums[i] = 0，则j保持不变，满足；
若nums[i] != 0，交换后j增一，仍然满足
*/
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i == j {
				j++
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
