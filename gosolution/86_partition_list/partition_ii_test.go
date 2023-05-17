package main

import (
	. "github.com/chinabobo/leetcode/model"
	"github.com/chinabobo/leetcode/util/linkedlist"
	"reflect"
	"testing"
)

func Test_partition2(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"case1",
			args{
				linkedlist.NewLinkedList([]int{1, 4, 3, 2, 5, 2, 3}),
				3,
			},
			linkedlist.NewLinkedList([]int{1, 2, 2, 3, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition2(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition2() = %v, want %v", got, tt.want)
			}
		})
	}
}
