package util

import (
	"reflect"
	. "github.com/chinabobo/leetcode/model"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}


func Are2DArraysEqual(arr1, arr2 [][]int) bool {
	// 检查数组长度是否相等
	if len(arr1) != len(arr2) {
		return false
	}

	// 逐个比较元素
	for i := 0; i < len(arr1); i++ {
		// 检查子数组长度是否相等
		if !reflect.DeepEqual(arr1[i], arr2[i]) {
			return false
		}
	}
	return true
}

func AreLinkedListsEqual(list1, list2 *ListNode) bool {
	// 循环遍历链表
	for list1 != nil && list2 != nil {
		// 如果节点值不相等，则链表不相等
		if list1.Val != list2.Val {
			return false
		}
		// 继续比较下一个节点
		list1 = list1.Next
		list2 = list2.Next
	}

	// 如果两个链表都遍历完了，且长度相等，则链表相等
	return list1 == nil && list2 == nil
}
