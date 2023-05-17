package main

import "testing"

func Test_countMustHave1In0Left(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{1},
			1,
		},
		{
			"case 2",
			args{2},
			2,
		},
		{
			"case 3",
			args{3},
			3,
		},
		{
			"case 4",
			args{4},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMustHave1In0Left(tt.args.n); got != tt.want {
				t.Errorf("countMustHave1In0Left() = %v, want %v", got, tt.want)
			}
		})
	}
}
