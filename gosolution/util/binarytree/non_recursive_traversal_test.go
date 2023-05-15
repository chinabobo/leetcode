package binarytree

import (
	"github.com/chinabobo/leetcode/gosolution/model"
	"reflect"
	"testing"
)

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
			"case1",
			args{
				BuildTree([]int{5, 4, 9, 2, 8, 1, 6, 7}),
			},
			[]int{1, 2, 4, 5, 6, 7, 8, 9},
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
			"case2",
			args{
				BuildTree([]int{5, 4, 9, 2, 8, 1, 6, 7}),
			},
			[]int{1, 2, 4, 7, 6, 8, 9, 5},
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

func Test_preorderTraversalNonRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case3",
			args{
				BuildTree([]int{5, 4, 9, 2, 8, 1, 6, 7}),
			},
			[]int{5, 4, 2, 1, 9, 8, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalNonRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalNonRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
