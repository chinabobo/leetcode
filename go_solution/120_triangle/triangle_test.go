package main

import "testing"

func Test_minimumTotalBottomUpDp(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minimum total bottom",
			args: args{
				triangle: [][]int{
					{2},
					{3, 4},
					{6, 5, 7},
					{4, 1, 8, 3},
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalBottomUpDp(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotalBottomUpDp() = %v, want %v", got, tt.want)
			}
		})
	}
}
