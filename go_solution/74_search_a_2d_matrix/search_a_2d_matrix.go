package main

import "fmt"

// https://leetcode.cn/problems/search-a-2d-matrix/

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。

func main() {
	aa := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	bb := [][]int{
		{1, 1},
	}
	cc := [][]int{
		{1},
		{3},
	}
	dd := [][]int{
		{1, 1},
		{2, 2},
	}
	fmt.Println(searchMatrix(aa, 3))
	fmt.Println(searchMatrix(bb, 1))
	fmt.Println(searchMatrix(cc, 0))
	fmt.Println(searchMatrix(dd, 0))
	fmt.Println(searchMatrix(dd, 3))
	fmt.Println(searchMatrix(aa, 30))
}

// 思路：将2纬数组转为1维数组 进行二分搜索

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		// 获取2纬数组对应值
		val := matrix[mid/col][mid%col]
		if val > target {
			end = mid
		} else if val < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}
	return false
}

func searchMatrixSelf(matrix [][]int, target int) bool {
	if len(matrix) == 1 {
		hasExist, _ := binarySearchSelf(matrix[0], target)
		if hasExist {
			return true
		} else {
			return false
		}
	}
	if len(matrix[0]) == 1 {
		heightSlice := make([]int, 0)

		for _, line := range matrix {
			heightSlice = append(heightSlice, line[0])
		}
		hasExist, _ := binarySearchSelf(heightSlice, target)
		if hasExist {
			return true
		} else {
			return false
		}
	}

	heightSlice := make([]int, 0)

	for _, line := range matrix {
		heightSlice = append(heightSlice, line[0])
	}

	hasExist, index := binarySearchSelf(heightSlice, target)
	if hasExist {
		return true
	} else {
		if index < 0 || index >= len(matrix) {
			return false
		}
		hasExist1, _ := binarySearchSelf(matrix[index], target)

		if hasExist1 {
			return true
		}
	}
	return false
}

func binarySearchSelf(sli []int, target int) (bool, int) {
	start := 0
	end := len(sli) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if sli[mid] < target {
			start = mid
		} else if sli[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if sli[start] == target {
		return true, start
	} else if sli[start] > target {
		return false, start - 1
	} else if sli[start] < target && sli[end] > target {
		return false, start
	} else if sli[end] == target {
		return true, end
	} else if sli[end] < target {
		return false, end
	}
	return false, 0
}
