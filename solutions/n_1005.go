package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	rocks := make([]int, n)
	var total int
	for i := 0; i < n; i++ {
		fmt.Scan(&rocks[i])
		total += rocks[i]
	}

	minDiff := total
	for mask := 0; mask < (1 << n); mask++ {
		var sum1 int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum1 += rocks[i]
			}
		}
		sum2 := total - sum1
		diff := abs(sum1 - sum2)
		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Println(minDiff)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
