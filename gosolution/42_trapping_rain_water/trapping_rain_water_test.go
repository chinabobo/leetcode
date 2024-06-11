package main

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.input); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
