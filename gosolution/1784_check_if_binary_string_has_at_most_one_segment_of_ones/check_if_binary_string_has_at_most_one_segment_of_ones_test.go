package main

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{"1001"},
			false,
		},
		{
			"case2",
			args{"110"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.args.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
