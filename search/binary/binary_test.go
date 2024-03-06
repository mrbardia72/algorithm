package main

import "testing"

func TestBinary(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found target",
			args: args{
				arr:    []int{1, 3, 5, 7, 9, 11},
				target: 7,
			},
			want: 3,
		},
		{
			name: "not found target",
			args: args{
				arr:    []int{1, 3, 5, 7, 9, 11},
				target: 17,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Binary(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("Binary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
