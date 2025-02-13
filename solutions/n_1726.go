package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&xs[i], &ys[i])
	}

	sort.Ints(xs)
	sort.Ints(ys)
	sum := 0

	for i := 0; i < n; i++ {
		sum += xs[i] * (2*i - n + 1)
		sum += ys[i] * (2*i - n + 1)
	}

	sum *= 2
	walks := n * (n - 1)

	res := sum / walks
	fmt.Println(res)
}
