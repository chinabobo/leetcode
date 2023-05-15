package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	表达式计算
给定一个字符虽形式的表达式，保证每个字符星表达式中仅包备加(+)这1种运算符，计着并输出零达式结黑要注意的是，+号两边的数据仅可能包含数字字符、小数点字符与持殊学符，特林字符包括 !, @, #，
这些特殊字将的加法运算有待别的规则
	! + ! = 0
	! + @ = 13
	! + # = 4
    @ + @ = 7
	@ + # = 20
	# + # = 5

注意：
1. 保证每个表达式仅包含一个运算符
2. 保证表达式一定可运算且有数据结果
3. 保证运算符两边数据有效(不会出现包含两个小数点之类的无效数据)
4. 表达式内不存在空格
5. 特殊字符的加法运算符合交换律
6. 如果表达式中包合特殊字符，则运算中不会出现数字与特殊字符的加法运算
7. 表达式两边的数据均不以0开头，比如不会出现这样的表达式:0250+0110

eg:
input :
	15
	123.45#1+126.53@

output :
	250.0001

explain:
	#+@-20，即进2位，表达式结果为250.0001，竖式运算如下:
	123.45#1
	126.53@
	---------
	250.0001

*/

var checkSymbols = make(map[string]bool)
var ansMap = make(map[[2]string]float64)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	// aa := "12345#+12653@"
	// bb := "1#3+1#0.4"
	// cc := "123.45#1+126.53@"

	symbols := []string{
		"!",
		"@",
		"#",
	}

	checkSymbols = map[string]bool{
		"!": true,
		"@": true,
		"#": true,
	}

	ansMap = map[[2]string]float64{
		{symbols[0], symbols[0]}: 0,
		{symbols[0], symbols[1]}: 13,
		{symbols[1], symbols[0]}: 13,
		{symbols[0], symbols[2]}: 4,
		{symbols[2], symbols[0]}: 4,
		{symbols[1], symbols[1]}: 7,
		{symbols[1], symbols[2]}: 20,
		{symbols[2], symbols[1]}: 20,
		{symbols[2], symbols[2]}: 5,
	}

	fmt.Println(calculator(s))
	// fmt.Println(calculator(bb))
	// fmt.Println(calculator(cc))
}

func calculator(src string) string {
	items := strings.Split(src, "+")
	var needAdd float64
	for _, v := range []rune(items[0]) {
		if checkSymbols[string(v)] {
			items2 := strings.Split(items[1], string(v))
			_, l := findLetter(items[1])
			hasD, index := findD(items[0])

			if hasD {
				needAdd = ansMap[[2]string{
					string(v),
					l,
				}] * math.Pow(10, float64(index-len(items2[0])+1))

			} else {
				needAdd = ansMap[[2]string{
					string(v),
					l,
				}] * math.Pow(10, float64(len(items[1])-len(items2[0])))
			}

			// fmt.Println(needAdd)
		}
	}
	res := dealSrcStr(items[0]) + dealSrcStr(items[1]) + needAdd
	return fmt.Sprintf("%g", res)
}

func findD(src string) (bool, int) {
	if len(src) == 0 {
		return false, 0
	}
	for i, v := range []rune(src) {
		if string(v) == "." {
			return true, i
		}
	}
	return false, 0
}

func dealSrcStr(src string) float64 {
	var it []string
	for _, v := range []rune(src) {
		if checkSymbols[string(v)] {
			it = strings.Split(src, string(v))
		}
	}
	res := it[0] + "0" + it[1]
	float, _ := strconv.ParseFloat(res, 64)
	return float
}

func findLetter(src string) (bool, string) {
	for _, v := range []rune(src) {
		if checkSymbols[string(v)] {
			return true, string(v)
		}
	}
	return false, ""
}
