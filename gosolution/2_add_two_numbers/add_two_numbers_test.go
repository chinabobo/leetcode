package main

import (
	. "github.com/chinabobo/leetcode/model"
	"github.com/chinabobo/leetcode/util"
	"github.com/chinabobo/leetcode/util/linkedlist"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		{
			name: "case1",
			args: args{
				linkedlist.NewLinkedList([]int{2, 4, 3}),
				linkedlist.NewLinkedList([]int{5, 6, 4}),
			},
			wantHead: linkedlist.NewLinkedList([]int{7, 0, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if gotHead := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(gotHead, tt.wantHead) {
			//	t.Errorf("addTwoNumbers() = %v, want %v", gotHead, tt.wantHead)
			//}
			if gotHead := addTwoNumbers(tt.args.l1, tt.args.l2); !util.AreLinkedListsEqual(gotHead, tt.wantHead) {
				t.Errorf("threeSum() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}
