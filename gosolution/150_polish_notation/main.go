package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/submissions/
// 思路：通过栈保存原来的元素，遇到表达式弹出运算，再推入结果，重复这个过程

func main() {
	// tokens := []string{"2","1","+","3","*"}
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err != nil {
			var res int
			switch v {
			case "*":
				res = stack[len(stack)-1] * stack[len(stack)-2]
			case "+":
				res = stack[len(stack)-1] + stack[len(stack)-2]
			case "-":
				res = stack[len(stack)-2] - stack[len(stack)-1]
			case "/":
				res = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
			fmt.Println(res)
			fmt.Println(stack)
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}
