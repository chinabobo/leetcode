package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.cn/problems/implement-rand10-using-rand7/

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(rand10())
	}
}

func rand10() int {
	for {
		x := (rand7()-1)*7 + (rand7() - 1) // 0 - 48
		// 0, 48
		if x >= 1 && x <= 40 {
			return x%10 + 1
		}
		// 41, 48
		x = (x%40)*7 + rand7() // 1 - 63
		if x <= 60 {
			return x%10 + 1
		}

		// 61, 62, 63
		x = (x-61)*7 + rand7() // 1 - 21
		if x <= 20 {
			return x%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
