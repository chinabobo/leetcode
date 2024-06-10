package main

import (
	"testing"
	. "github.com/chinabobo/leetcode/util"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case1",
			args{
				nums: []int{-1,0,1,2,-1,-4},
			},
			[][]int{
				{-1,-1,2},
				{-1,0,1},
			},
		},
		{
			"case2",
			args{
				nums: []int{0,1,1},
			},
			nil,
		},
		{
			"case3",
			args{
				nums: []int{0,0,0},
			},
			[][]int{
				{0,0,0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !Are2DArraysEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
