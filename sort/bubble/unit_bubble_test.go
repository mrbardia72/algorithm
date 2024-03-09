package main

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort array",
			args: args{arr: []int{15, 12, 60, 59, 2, 1}},
			want: []int{1, 2, 12, 15, 59, 60},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
