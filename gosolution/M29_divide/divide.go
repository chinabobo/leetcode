package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2147483647
	b := 3
	res := divide(a, b)
	fmt.Println("start ")
	fmt.Println(res)
	fmt.Println(a / b)

}

func divide(a int, b int) int {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	if a == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	aa := a
	if aa < 0 {
		aa = 0 - a
	}
	bb := b
	if bb < 0 {
		bb = 0 - b
	}

	var t, tt, res int
	t = aa
	for i := 1; ; i++ {
		t = t - bb
		if t <= 0 {
			tt = i
			break
		}
	}
	if t == 0 {
		res = tt
	} else if t < 0 {
		res = tt - 1
	}
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		res = 0 - res
	}
	if res > 2147483647 || res < -2147483648 {
		return 2147483647
	}
	return res
}
