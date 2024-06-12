package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case1",
			args{
				input: []int{1, 2, 3},
			},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			"case2",
			args{
				input: []int{1, 0},
			},
			[][]int{
				{1,0},
				{0,1},
			},
		},
		{
			"case3",
			args{
				input: []int{1},
			},
			[][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.input); !are2DArraysEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to convert a 2D array to a map of 1D arrays
func convertToSet(arr [][]int) map[string]struct{} {
	set := make(map[string]struct{})
	for _, row := range arr {
		key := fmt.Sprint(row)
		set[key] = struct{}{}
	}
	return set
}

// Helper function to check if two 2D arrays have the same elements
func are2DArraysEqual(arr1, arr2 [][]int) bool {
	set1 := convertToSet(arr1)
	set2 := convertToSet(arr2)
	return reflect.DeepEqual(set1, set2)
}