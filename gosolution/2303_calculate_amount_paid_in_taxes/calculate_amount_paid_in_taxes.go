package main

import "fmt"

/*
https://leetcode.cn/problems/calculate-amount-paid-in-taxes/
*/
func main() {
	aa := [][]int{
		{3, 50},
		{7, 10},
		{12, 25},
	}
	fmt.Println(calculateTax(aa, 10))
}

func calculateTax(brackets [][]int, income int) float64 {
	var res int
	upper := 0
	lower := 0
	for i := range brackets {
		upper = brackets[i][0]
		res += (min(income, upper) - lower) * brackets[i][1]
		if upper >= income {
			break
		}
		lower = upper
	}
	return float64(res) / 100
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func calculateTaxSelf(brackets [][]int, income int) float64 {
	if income == 0 {
		return 0
	}
	var res float64
	res = 0
	for i, v := range brackets {
		if income <= v[0] {
			if i != 0 {
				res += float64(income-brackets[i-1][0]) * float64(v[1]) / 100
			} else {
				res += float64(income) * float64(v[1]) / 100
			}
			break
		}
		if i != 0 {
			res += float64(v[0]-brackets[i-1][0]) * float64(v[1]) / 100
		} else {
			res += float64(v[0]) * float64(v[1]) / 100
		}
		fmt.Println(res)
	}
	return res
}
