package main

import "testing"

func TestIsEven(t *testing.T) {
	if isEven(0) != true {
		t.Errorf("Expected to be true for number 0, got %v", isEven(0))
	}

	if isEven(100) != true {
		t.Errorf("Expected to be true for number 100, got %v", isEven(100))
	}

	if isEven(1) != false {
		t.Errorf("Expected to be false for number 1, got %v", isEven(1))
	}

	if isEven(99) != false {
		t.Errorf("Expected to be false for number 99, got %v", isEven(99))
	}
}
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
