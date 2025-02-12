package main

import (
	"fmt"
	"sort"
)

type charI struct {
	letter rune
	index  int
}

// НЕ УКЛАДЫВАЕТСЯ ПО ВРЕМЕНИ
func main() {
	var n int
	var input string
	fmt.Scan(&n)
	fmt.Scan(&input)

	n--

	var s []charI
	for i, c := range input {
		s = append(s, charI{letter: c, index: i})
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].letter < s[j].letter
	})

	j := n
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[j].letter))
		j = s[j].index
	}
}
