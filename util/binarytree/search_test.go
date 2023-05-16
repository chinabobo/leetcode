package binarytree

import (
	. "github.com/chinabobo/leetcode/model"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case-levelOrder",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversalBottomup(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case-preorderTraversalBottomup",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalBottomup(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalBottomup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversalTopdown(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case-preorderTraversalTopdown",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalTopdown(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalTopdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
