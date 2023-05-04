package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	给定 n 个不同元素的数组，设计算法等概率取 m 个不同的元素
*/

func main() {
	aa := []int{
		1, 2, 3, 4, 5, 6,
	}
	m := 3
	fmt.Println(randomM(aa, m))
}

func randomM(src []int, m int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(src), func(i, j int) {
		src[i], src[j] = src[j], src[i]
	})
	return src[:m]
}
