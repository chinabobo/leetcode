package binarytree

import (
	. "github.com/chinabobo/leetcode/model"
	"reflect"
	"testing"
)

func Test_preorderTraversalRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
		res  *[]int
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
				&[]int{},
			},
			[]int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversalRecursive(tt.args.root, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversalRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
		res  *[]int
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
				&[]int{},
			},
			[]int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalRecursive(tt.args.root, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversalRecursive(t *testing.T) {
	type args struct {
		root *TreeNode
		res  *[]int
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
				&[]int{},
			},
			[]int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversalRecursive(tt.args.root, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
