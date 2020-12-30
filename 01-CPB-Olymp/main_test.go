package main

import "testing"

func Test_fn(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		m       int
		wantRes int
	}{
		{"1", 5, 3, 6},
		{"2", 4, 3, 7},
		{"3", 2, 2, 4},
		{"4", 1, 0, 1},
		{"5", 3, 7, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := fn(tt.n, tt.m); gotRes != tt.wantRes {
				t.Errorf("fn() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
