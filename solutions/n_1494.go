package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	balls := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&balls[i])
	}

	stack := make([]int, 0)
	next := 1

	for _, ball := range balls {
		for next <= n && (len(stack) == 0 || stack[len(stack)-1] != ball) {
			stack = append(stack, next)
			next++
		}
		if len(stack) > 0 && stack[len(stack)-1] == ball {
			stack = stack[:len(stack)-1]
		} else {
			fmt.Println("Cheater")
			return
		}
	}
	fmt.Println("Not a proof")
}
