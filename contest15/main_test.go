package main

import "testing"

func Test_fn(t *testing.T) {
	tests := []struct {
		name    string
		k       int64
		a       []int64
		wantRes int64
	}{
		{"-4", 56, []int64{2, 1, 4, 3, 2}, 5},
		{"0", 10, []int64{100, 11}, 10},
		{"1", 1, []int64{20100, 1, 202, 202}, 99},
		{"2", 100, []int64{1, 1, 1}, 0},
		{"-3", 100, []int64{1, 1, 2, 4, 9}, 1},
		{"-2", 100, []int64{1, 2, 2, 4, 8}, 1},
		{"-1", 100, []int64{1, 1, 2, 4, 8}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := fn(tt.k, tt.a); gotRes != tt.wantRes {
				t.Errorf("fn() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
func fn(k int64, a []int64) int64 {
	x := int64(0)
	sum := a[0]
	for i := 1; i < len(a); i++ {
		pr := (100*a[i] - k*sum + k - 1) / k
		x = max(x, pr)
		sum += a[i]
	}
	return x
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
