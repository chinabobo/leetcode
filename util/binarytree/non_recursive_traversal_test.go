package binarytree

import (
	. "github.com/chinabobo/leetcode/model"
	"reflect"
	"testing"
)

func TestPreorderTraversalNonRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case-1",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreorderTraversalNonRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversalNonRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversalNonRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case-1",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[]int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalNonRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalNonRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversalNonRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case-1",
			args{
				BuildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			[]int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversalNonRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalNonRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
