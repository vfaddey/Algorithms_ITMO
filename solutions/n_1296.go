package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println(0)
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	maxSum, currentSum := 0, 0
	for i := 0; i < n; i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	fmt.Println(maxSum)
}
