package main

import "testing"

func Test_calculateMaxValue(t *testing.T) {
	type args struct {
		n int
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case-1",
			args{
				6,
				"010101",
			},
			7,
		},
		{
			"case-2",
			args{
				20,
				"11111000111011101100",
			},
			94,
		},
		{
			"case-3",
			args{
				4,
				"1100",
			},
			6,
		},
		{
			"case-4",
			args{
				3,
				"111",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMaxValue(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("calculateMaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
