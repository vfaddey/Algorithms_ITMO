package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var m, n, k int
	fmt.Fscan(in, &m, &n, &k)

	rowBlacks := make([][]int, m+1)
	colBlacks := make([][]int, n+1)

	for t := 0; t < k; t++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		rowBlacks[x] = append(rowBlacks[x], y)
		colBlacks[y] = append(colBlacks[y], x)
	}

	for i := 1; i <= m; i++ {
		sort.Ints(rowBlacks[i])
	}
	for j := 1; j <= n; j++ {
		sort.Ints(colBlacks[j])
	}

	stripesH := 0
	hSingles := make(map[uint64]struct{})

	for i := 1; i <= m; i++ {
		prev := 0
		for _, y := range rowBlacks[i] {
			length := y - prev - 1
			if length >= 2 {
				stripesH++
			} else if length == 1 {
				key := uint64(i)<<32 | uint64(prev+1)
				hSingles[key] = struct{}{}
			}
			prev = y
		}
		length := n - prev
		if length >= 2 {
			stripesH++
		} else if length == 1 {
			key := uint64(i)<<32 | uint64(prev+1)
			hSingles[key] = struct{}{}
		}
	}
	stripesV := 0
	isolated := 0

	// Scan each column for white runs
	for j := 1; j <= n; j++ {
		prev := 0
		for _, x := range colBlacks[j] {
			length := x - prev - 1
			if length >= 2 {
				stripesV++
			} else if length == 1 {
				key := uint64(prev+1)<<32 | uint64(j)
				if _, ok := hSingles[key]; ok {
					isolated++
				}
			}
			prev = x
		}
		length := m - prev
		if length >= 2 {
			stripesV++
		} else if length == 1 {
			key := uint64(prev+1)<<32 | uint64(j)
			if _, ok := hSingles[key]; ok {
				isolated++
			}
		}
	}

	fmt.Println(stripesH + stripesV + isolated)
}
