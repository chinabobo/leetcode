package main

import "testing"

func Test_maxArea(t *testing.T) {
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
				input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.input); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*  use code runner plugin in vscode

in *.go use keyboard shortcut Ctrl+Alt+K to run *_test.go

"code-runner.customCommand": "cd $dir && go test $fileNameWithoutExt_test.go $fileName"

*/
