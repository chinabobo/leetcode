package main

import "testing"

func Test_backpack(t *testing.T) {
	type args struct {
		m int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				12,
				[]int{2, 3, 5, 7},
			},
			12,
		},
		{
			"case2",
			args{
				10,
				[]int{3, 4, 8, 5},
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backpack(tt.args.m, tt.args.a); got != tt.want {
				t.Errorf("backpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
