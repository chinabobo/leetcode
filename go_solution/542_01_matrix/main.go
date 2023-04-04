package main

import "fmt"

// BFS 从0进队列，弹出之后计算上下左右的结果，将上下左右重新进队列进行二层操作
// 0 0 0 0
// 0 x 0 0
// x x x 0
// 0 x 0 0

// 0 0 0 0
// 0 1 0 0
// 1 x 1 0
// 0 1 0 0

// 0 0 0 0
// 0 1 0 0
// 1 2 1 0
// 0 1 0 0

// https://leetcode.cn/problems/01-matrix/

func main() {
	input := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	updateMatrix(input)
	// fmt.Println(updateMatrix(input))
}
func updateMatrix(matrix [][]int) [][]int {
	q := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// 进队列
				point := []int{i, j}
				q = append(q, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(q) != 0 {

		// 出队列
		point := q[0]
		q = q[1:]
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			fmt.Println("point : ", point)
			fmt.Printf("x : %d, y : %d \n", x, y)

			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				fmt.Printf("[BFS] x : %d, y : %d \n", x, y)

				matrix[x][y] = matrix[point[0]][point[1]] + 1
				// 将当前的元素进队列，进行一次BFS
				q = append(q, []int{x, y})
				fmt.Println(q)

				for _, vv := range matrix {
					fmt.Println(vv)
				}
			}
		}
	}
	return matrix

}
