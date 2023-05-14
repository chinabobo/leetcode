package main

import (
	"fmt"
	"math"
)

/*
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

*/
func main() {
	aa := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(aa))
}

/*
记录每次的min价格，和最大利润，

maxprofit = max(profit, prices[i]-min)
*/
func maxProfit(prices []int) int {
	var profit int
	min := math.MaxInt32
	if len(prices) == 1 {
		return 0
	}
	for i := 0; i < len(prices); i++ {
		min = minf(min, prices[i])
		if i-1 >= 0 && prices[i] >= prices[i-1] {
			profit = maxf(profit, prices[i]-min)
		}

	}
	return profit
}

func minf(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
