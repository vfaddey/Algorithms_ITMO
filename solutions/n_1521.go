package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	r := ring.New(n)
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}
	for r.Len() > 1 {
		r = r.Move(k - 1)
		fmt.Println(r.Value)
		next := r.Next()
		r.Prev().Unlink(1)
		r = next
	}
	fmt.Println(r.Value)
}
