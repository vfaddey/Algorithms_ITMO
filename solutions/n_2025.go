package main

import "fmt"

func main() {
	var t, n, k int
	if _, err := fmt.Scan(&t); err != nil {
		return
	}
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &k)
		res := solve(n, k)
		fmt.Println(res)
	}
}

func solve(n int, k int) int {
	eq := n / k
	rest := n % k
	teams := make([]int, k)
	for i := 0; i < k; i++ {
		if rest > 0 {
			teams[i] = eq + 1
			rest--
		} else {
			teams[i] = eq
		}
	}
	var total int
	for _, team := range teams {
		total += team * (n - team)
	}

	return total / 2
}
