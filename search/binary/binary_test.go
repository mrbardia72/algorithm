package main

import "testing"

func TestBinary(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			// TODO: Add test cases.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binary(tt.args.arr, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Binary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
