package main

import "testing"

func Test_backpackII(t *testing.T) {
	type args struct {
		m int
		a []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				10,
				[]int{2, 3, 5, 7},
				[]int{1, 5, 2, 4},
			},
			9,
		},
		{
			"case2",
			args{
				10,
				[]int{2, 3, 8},
				[]int{2, 5, 8},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backpackII(tt.args.m, tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("backpackII() = %v, want %v", got, tt.want)
			}
		})
	}
}
