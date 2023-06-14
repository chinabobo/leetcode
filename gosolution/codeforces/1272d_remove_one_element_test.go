package main

import "testing"

func Test_find(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{1, 2, 5, 3, 4},
			},
			4,
		},
		{
			"case2",
			args{
				[]int{1, 2},
			},
			2,
		},
		{
			"case3",
			args{
				[]int{6, 5, 4, 3, 2, 4, 3},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLen(tt.args.src); got != tt.want {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
