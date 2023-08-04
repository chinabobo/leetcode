package main

import "testing"

func Test_twoEggDrop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{100},
			14,
		},
		{
			"case2",
			args{2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDrop(tt.args.n); got != tt.want {
				t.Errorf("twoEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoEggDropOpt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{100},
			14,
		},
		{
			"case2",
			args{2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDropOpt(tt.args.n); got != tt.want {
				t.Errorf("twoEggDropOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}
