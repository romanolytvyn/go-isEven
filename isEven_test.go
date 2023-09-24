package main

import "testing"

func Test_isEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "2", args: args{num: 2}, want: true},
		{name: "13", args: args{num: 13}, want: false},
		{name: "42", args: args{num: 42}, want: true},
		{name: "69", args: args{num: 69}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.num); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
