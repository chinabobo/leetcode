package main

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{3, 4, 7, 2, -3, 1, 4, 2},
				7,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraySumBruteForceSolution(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{3, 4, 7, 2, -3, 1, 4, 2},
				7,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumBruteForceSolution(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumBruteForceSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraySumSelf(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{3, 4, 7, 2, -3, 1, 4, 2},
				7,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySumSelf(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySumSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
