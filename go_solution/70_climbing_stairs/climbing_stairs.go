package main

import "fmt"

/*

f(i) = f(i-1) + f(i-2)

*/

func main() {
	fmt.Println(climbStairs(3)) // 3
}

func climbStairs(n int) int {
	var a, b int
	a = 1
	b = 1

	for i := 1; i < n; i++ {
		t := a + b
		a = b
		b = t
	}
	return b
}
