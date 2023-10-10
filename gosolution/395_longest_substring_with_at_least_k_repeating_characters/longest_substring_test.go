package main

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "case 1",
			args: args{
				s: "aaabb",
				k: 3,
			},
			wantAns: 3,
		},
		{
			name: "case 2",
			args: args{
				s: "ababbc",
				k: 2,
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubstring(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("longestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
