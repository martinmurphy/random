package main

import (
	"testing"
)

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a<b", args{10, 20}, 10},
		{"a=b", args{20, 20}, 20},
		{"a>b", args{100, 20}, 20},
		{"-a<b", args{-10, 20}, -10},
		{"-a<-b", args{-20, -10}, -20},
		{"-a=-b", args{-10, -10}, -10},
		{"-a>-b", args{-10, -20}, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("\"%v\": min() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		v int
		s []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"value in set", args{10, []int{5, 10, 20, 30}}, true},
		{"value not in set", args{1, []int{5, 10, 20, 30}}, false},
		{"empty set", args{10, []int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.v, tt.args.s); got != tt.want {
				t.Errorf("\"%v\": contains() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
