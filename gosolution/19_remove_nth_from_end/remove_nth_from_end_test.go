package removenthfromend

import (
	"github.com/chinabobo/leetcode/model"
	"github.com/chinabobo/leetcode/util"
	"github.com/chinabobo/leetcode/util/linkedlist"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		input *model.ListNode
		n int
	}
	tests := []struct {
		name string
		args args
		want *model.ListNode
	}{
		{
			"case1",
			args{
				input: linkedlist.NewLinkedList([]int{1, 2, 3, 4, 5}),
				n: 2,
			},
			linkedlist.NewLinkedList([]int{1, 2, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.input, tt.args.n); !util.AreLinkedListsEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}