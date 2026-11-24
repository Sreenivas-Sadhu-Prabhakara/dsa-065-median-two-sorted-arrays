package main

import (
	"fmt"
	"math"
)

// Median returns the median of two individually sorted arrays in
// O(log min(n,m)), by binary searching the partition of the shorter array.
func Median(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	n, m := len(a), len(b)
	if n+m == 0 {
		return 0
	}
	half := (n + m + 1) / 2
	lo, hi := 0, n
	for lo <= hi {
		i := (lo + hi) / 2 // elements taken from a's left side
		j := half - i      // the rest come from b
		aLeft, aRight := math.MinInt, math.MaxInt
		if i > 0 {
			aLeft = a[i-1]
		}
		if i < n {
			aRight = a[i]
		}
		bLeft, bRight := math.MinInt, math.MaxInt
		if j > 0 {
			bLeft = b[j-1]
		}
		if j < m {
			bRight = b[j]
		}
		switch {
		case aLeft <= bRight && bLeft <= aRight:
			if (n+m)%2 == 1 {
				return float64(max(aLeft, bLeft))
			}
			return (float64(max(aLeft, bLeft)) + float64(min(aRight, bRight))) / 2
		case aLeft > bRight:
			hi = i - 1
		default:
			lo = i + 1
		}
	}
	return 0 // unreachable for valid sorted inputs
}

func main() {
	fmt.Println(Median([]int{1, 3}, []int{2}))
	fmt.Println(Median([]int{1, 2}, []int{3, 4}))
}
