package main

import "fmt"

/*
Problem:
	Allocating resource
	给定一个管理ID的资源池，可以从资源池中分配资源ID和释放资源ID
	分配方式有动态分配和指定分配，动态分配是从资源池的开始分配一个资源ID
	指定分配是指定一个资源ID进行分配，无论哪种分配方式释放资源ID时都第要放到资源池的尾部。
	执行一系列操作后，请问资源池的第一个空闲资源ID应该是多少?
	注意
	资源池的初始顺序是从小到大
	资源池中空闲资源ID不足时，动态分配失败，对资源池不进行低何择作指定分配资源ID巴经披占用或者不在资源池范围内时，对资源池不进行任何操作。
	释放资源ID不在资源池范围内时或者已经是空闲资源ID时，对资源池不进行任何操作，保证每个用例最启都有空闲资源ID.
	输入
	第一行是资源池的范围，第二行是操作个数。
	第三行开始，第一个数字代表操作类型
	1表示动态分配，2表示指定分配，3表示释放
	如果第一个数字是1，第二个表示分配的个数
	如果第一个教字是2，第二个表示分配的资源ID
	如果第一个数字是3，第二个表示释放的资源ID

	input:
	1 3
	2
	1 1
	3 1

	output:
	2

	explain:
	第一行资源池范围是[1 3]，资源池的初始顺序是1->2->3
	第二行操作个数有2个
	第三行动态分配1个资源ID，资源池中剩余的资源ID顺序是2->3
	第四行释放1个资源ID，资源ID是1，资源池中剩余的盗源ID顺序是2->3
	执行以上操作后，资源池的第一个空闲资源ID是2

*/

// solution: use stack

func main() {
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
	var count int
	_, err = fmt.Scan(&count)
	if err != nil {
		return
	}
	fmt.Println(count)

	operation := make([][]int, count)
	for i := 0; i < count; i++ {
		operation[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			var v int
			fmt.Scan(&v)
			operation[i][j] = v
		}
	}
	fmt.Println(operation)

	// m := 1
	// n := 3
	// count := 2
	// operation := [][]int{
	// 	{1, 1},
	// 	{3, 1},
	// }
	fmt.Println(allocatingResource(m, n, operation))

}

func allocatingResource(m, n int, operation [][]int) int {
	aa := make([]int, n-m+1)
	for i := range aa {
		aa[i] = m + i
	}
	for _, v := range operation {
		switch v[0] {
		case 1:
			aa = aa[1:]
		case 2:
			for i, vv := range aa {
				if vv == v[1] {
					aa = append(aa[:i], aa[i+1:]...)
				}
			}
		case 3:
			aa = append(aa, v[1])
		}
	}
	return aa[0]
}
