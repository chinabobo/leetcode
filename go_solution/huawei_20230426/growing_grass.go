package main

import (
	"fmt"
	"math"
)

/*
Problem:
	Growing grass
	将n种不同的草随机种在一块广道无垠的二维平面上(直角坐标系内)
	给定二维数组 points 表示第0天所有的初始位置
	每天，被草覆盖的点会同外梦延到它的8个邻居点。注意，初始状态下，可能都会种草在同一点上
	现给定一个整数M，问最少需要多少天，方能找到一点同时至少有M种草?

	输入
	第一行输入整数M.(2 <= M <= n)
	第二行输入草的种数n(2 <= n <= 50)
	后面连续n行输入草初始位置[Xi, Yi] (1 <= Xi, Yi <= 10^9)

	input:
	2
	2
	2 1
	6 2

	output:
	2

	explain:

			  Y
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 4 |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 3 |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 2 |   |   |   |   |   | B |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 1 |   | A |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 0 | 1 | 2 | 3 | 4 | 5 | 6 |   |   |   | X
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+

			  Y
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 4 |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 3 |   |   |   |   | B | B | B |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 2 | A | A | A |   | B | B | B |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 1 | A | A | A |   | B | B | B |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 0 | A | A | A | 4 | 5 | 6 |   |   |   | X
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+

			  Y
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | 4 |   |   |   | B | B | B | B | B |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | A | A | A | A | * | B | B | B | B |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | A | A | A | A | * | B | B | B | B |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | A | A | A | A | * | B | B | B | B |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | A | A | A | A | * | B | B | B | B |   | X
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   | A | A | A | A | A |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+
	|   |   |   |   |   |   |   |   |   |   |   |   |
	+---+---+---+---+---+---+---+---+---+---+---+---+

	在第2天 点(4, 0), (4, 1), (4, 2), (4, 3) 将有2种草

	eg2:
	input:
	2
	3
	2 1
	6 2
	100 100

	output:
	2

Solution:
	找到所有点中某两个点的最小距离
	距离的定义是max((X1 - X2), (Y1 - Y2))

*/

// todo: 只是解决了m=2的情况

func main() {
	// aa := [][2]int{
	// 	{2, 1},
	// 	{6, 2},
	// 	{100, 100},
	// }
	var m int
	_, err := fmt.Scan(&m)
	if err != nil {
		return
	}
	var n int
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}
	aa := make([][2]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			var v int
			fmt.Scan(&v)
			aa[i][j] = v
		}
	}

	fmt.Println(aa)
	fmt.Println(findMofGrass(aa))
}
func findMofGrass(src [][2]int) int {
	minDistance := math.MaxInt
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if calculateDistance(src[i], src[j]) < minDistance {
				minDistance = calculateDistance(src[i], src[j])
			}
		}

	}
	return minDistance / 2
}

func calculateDistance(a, b [2]int) int {
	return max(differenceAbs(a[0], b[0]), differenceAbs(a[1], b[1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func differenceAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}