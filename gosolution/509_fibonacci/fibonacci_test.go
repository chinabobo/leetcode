package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				input: 2,
			},
			1,
		},
		{
			"case2",
			args{
				input: 3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.input); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
