package main

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {
	cases := []struct {
		a, b []int
		want float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1}, 1},
		{[]int{2, 2}, []int{2, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, []int{}, 3},
		{[]int{1, 2}, []int{-5, -4, -3}, -3},
		{[]int{100}, []int{1, 2, 3}, 2.5},
	}
	for _, c := range cases {
		if got := Median(c.a, c.b); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Median(%v,%v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
