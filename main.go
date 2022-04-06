package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type Dataset struct {
	name    string
	version string
}

var a []int = []int{1, 2, 3, 4, 5, 6, 7}
var rawLen = 7
var combineLen = 5

type UnixTime struct {
	time.Time
}

func main() {
	aa := []string{"a", "b", "c"}
	fmt.Println(aa)
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&aa)))

	b := append(aa[:1], aa[2:]...)

	// b := make([]string, 2)
	// at := copy(b, aa[0:1])
	// copy(b[at:], aa[2:3])

	fmt.Println(b)
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&b)))
	fmt.Println(aa)
	fmt.Println(*(*reflect.SliceHeader)(unsafe.Pointer(&aa)))

	// fmt.Println(b)
}

func combineloop(arr []string, r []string, i int, n int, output chan<- []string) {
	if n <= 0 {
		return
	}
	rlen := len(r) - n
	alen := len(arr)
	for j := i; j < alen; j++ {
		r[rlen] = arr[j]
		if n == 1 {
			or := make([]string, len(r))
			copy(or, r)
			output <- or
		} else {
			combineloop(arr, r, j+1, n-1, output)
		}
	}
}

func GetCombineMatch(arr []string, n int) []string {
	var matchArr []string
	output := make(chan []string)
	r := make([]string, n)
	go func() {
		combineloop(arr, r, 0, n, output)
		close(output)
	}()
	for arr := range output {
		var str string
		for i, each := range arr {
			if i == 0 {
				str = each
			} else {
				str = fmt.Sprintf("%v,%v", str, each)
				fmt.Println(str)
			}
		}
		matchArr = append(matchArr, str)
		fmt.Println(arr)
	}
	return matchArr
}

func combine() {
	fmt.Println("start combine")
	for i := 0; i <= 2; i++ {
		result := make([]int, combineLen)
		result[0] = a[i]
		// fmt.Println(i, arrLen)
		doProcess(result, i, 1)
	}
}

func doProcess(result []int, rawIndex int, curIndex int) {
	var choice int = rawLen - rawIndex + curIndex - combineLen
	// fmt.Printf("Choice: %d, rawLen: %d, rawIndex : %d, curIndex : %d \r\n", choice, rawLen, rawIndex, curIndex)
	var tResult []int
	for i := 0; i < choice; i++ {
		if i != 0 {
			tResult := make([]int, combineLen)
			copyArr(result, tResult)
		} else {
			tResult = result
		}
		// fmt.Println(curIndex)
		tResult[curIndex] = a[i+1+rawIndex]

		if curIndex+1 == combineLen {
			PrintIntArr(tResult)
			continue
		} else {
			doProcess(tResult, rawIndex+i+1, curIndex+1)
		}
	}
}
func copyArr(rawArr []int, target []int) {
	for i := 0; i < len(rawArr); i++ {
		target[i] = rawArr[i]
	}
}

func PrintIntArr(arr []int) {
	valuesText := []string{}
	for i := range arr {
		number := arr[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	fmt.Println(strings.Join(valuesText, ","))
}

func modUrl(rawUrl string) int {
	for _, item := range strings.Split(rawUrl, "/") {
		if IsNum(item) {
			id, _ := strconv.Atoi(item)
			return id
		}
	}
	return 0
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func sublist(a []int64, b []int64) []int64 {
	var c []int64
	temp := map[int64]struct{}{} // map[string]struct{}{}创建了一个key类型为String值类型为空struct的map，Equal -> make(map[string]struct{})

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}

func addlist(a []int64, b []int64) []int64 {
	var c []int64
	c = b
	temp := map[int64]struct{}{} // 创建一个key类型为String值类型为空struct的map  Equal -> make(map[int64]struct{})
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}

func GetHighestVersion() (string, error) {
	var versionList []float64
	publishedDatasets := []Dataset{
		{
			"a",
			"0.1",
		},
		{
			"b",
			"0.2",
		},
	}
	for _, dataset := range publishedDatasets {
		version, err := strconv.ParseFloat(dataset.version, 64)
		if err != nil {
			return "", err
		}
		versionList = append(versionList, version)
	}
	for i, a := range versionList {
		fmt.Println("[bb] versionList ", i, ":", a)
	}
	largestVersion := getMaxFloat(versionList)
	largestVersionStr := fmt.Sprintf("%.1f", largestVersion)
	return largestVersionStr, nil
}
func getMaxFloat(ary []float64) float64 {
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal < ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal
}
