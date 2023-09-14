package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "198"
	b := "292"
	fmt.Println(addBigData(a, b))
}

func addBigData(str1, str2 string) string {
	cnt1 := len(str1)
	cnt2 := len(str2)
	add := 0
	ans := ""
	for cnt1 > 0 || cnt2 > 0 {
		x := 0
		if cnt1 > 0 {
			x = int(str1[cnt1-1] - '0')
		}
		y := 0
		if cnt2 > 0 {
			y = int(str2[cnt2-1] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
		cnt2--
		cnt1--

	}
	return ans
}
