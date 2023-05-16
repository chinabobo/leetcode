package main

import (
	"fmt"
	"math"
	"strings"
)

/*
https://leetcode.cn/problems/fraction-to-recurring-decimal/

给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
输入：numerator = 4, denominator = 333
输出："0.(012)"

reference https://leetcode.cn/problems/fraction-to-recurring-decimal/solution/suan-fa-bi-ji-ha-xi-biao-shu-shi-chu-fa-eyj4l/
*/

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 { // 整除直接返回
		return fmt.Sprintf("%v", numerator/denominator)
	}

	sb := strings.Builder{}

	if numerator*denominator < 0 { // 遇到负数特殊处理
		sb.WriteString("-")
		numerator = int(math.Abs(float64(numerator)))
		denominator = int(math.Abs(float64(denominator)))
	}

	sb.WriteString(fmt.Sprintf("%v.", numerator/denominator)) // 遇到无法整除的情况，先将非小数部分写入结果

	numerator = numerator % denominator // 被除数取余，继续运算(竖式除法的核心)
	m := make(map[int]int)
	m[numerator] = sb.Len() // 在map中记录当前余数a第一次出现时结果集的长度
	for numerator != 0 {
		numerator *= 10 // 余数乘10继续除
		sb.WriteString(fmt.Sprintf("%v", numerator/denominator))
		numerator = numerator % denominator // 再次取余
		if m[numerator] != 0 {              // 判断新的余数之前是否出现过，如果出现过则将上次出现的位置前的字符串内容和两次出现中间的内容加上括号拼接在一起
			return fmt.Sprintf("%v(%v)", sb.String()[:m[numerator]], sb.String()[m[numerator]:sb.Len()])
		}
		m[numerator] = sb.Len() // 在map中记录当前余数a第一次出现时结果集的长度
	}
	return sb.String()
}
