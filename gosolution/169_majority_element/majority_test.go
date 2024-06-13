package main

import "testing"

func Test_majorityElement(t *testing.T) {
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
				input: []int{3, 2, 3},
			},
			3,
		},
		{
			"case2",
			args{
				input: []int{2, 2, 1, 1, 1, 2, 2},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.input); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}