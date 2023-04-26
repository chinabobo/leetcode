package main

import (
	"fmt"
)

/*
Problem:
	Find dependencies
	某部门在开发一个代码分析工具，需要分析代码模块之间的依赖关系，用来确定模块的初始化顺序，是否有循环依赖等问题。"批量初始化”"是指一次可以初始化一个或多个模块。
	例如模块1依赖模块2，模块3也依赖模块2，但模块1和3没有依赖关系。则必须先”批量初始化“模块2，再”批量初始化"模块1和3。现给定一组模块间的依赖关系，请计算需要”批量初始化”的次数。
	输入
	(1)第1行只有一个数字表示模块总数N。
	(2)随后的N行依次表示模块1到N的依赖数据。每行的第1个数据表示依赖的模块数量(不会超过N)，之后的数字表示当前模块依赖的模块ID序列，该序列不会重复出现相同的数字,模块ID的取值一定在[1M]之内。
	(3)模块总数N取值范围 1<=N<=1000
	(4)每一行里面的数字按1个空分隔

	input:
	5
	3 2 3 4
	1 5
	1 5
	1 5
	0

	output:
	3

	explain:
	共5个模块。
	模块1依赖模块2、3、4
	模块2赖模块5:模块3依赖模块5
	模块4依赖模块5
	模块5没有依赖任何模块
	批量初始化顺序为
	{5} -> {2, 3, 4} -> {1}
	共需”批量初始化”3次

*/

// solution: dfs

func main() {
	// aa := [][]int{
	// 	{2, 3, 4},
	// 	{5},
	// 	{5},
	// 	{5},
	// 	{},
	// }
	// bb := [][]int{
	// 	{2},
	// 	{3},
	// 	{1},
	// }
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	aa := make([][]int, n)

	var nn int
	for j := 0; j < n; j++ {
		fmt.Scan(&nn)
		aa[j] = make([]int, nn)
		for i := 0; i < nn; i++ {
			var b int
			fmt.Scan(&b)
			aa[j][i] = b
		}

	}
	fmt.Println(aa)

	fmt.Println(countDependenceInit(aa))
	// fmt.Println(countDependenceInit(bb))
}

func countDependenceInit(src [][]int) int {
	return dfs(src, 1, 0)
}

func dfs(src [][]int, id int, count int) int {
	if count > len(src) {
		return -1
	}
	count++
	if len(src[id-1]) == 0 {
		return count
	} else {
		return dfs(src, src[id-1][0], count)
	}
}
