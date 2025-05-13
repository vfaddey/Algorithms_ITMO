package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const PI = 3.14159265

type Pumpkin struct {
	Length  float64
	Degrees float64
	Index   int
}

func (p *Pumpkin) Create(fx, fy, idx int, r *bufio.Reader) {
	var x, y int
	fmt.Fscan(r, &x, &y)
	dx := float64(x - fx)
	dy := float64(y - fy)
	p.Index = idx + 1
	p.Length = dx*dx + dy*dy
	p.Degrees = math.Atan2(dy, dx) * 180.0 / PI
	if p.Degrees < 0 {
		p.Degrees += 360.0
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, fx, fy int
	fmt.Fscan(in, &n, &fx, &fy)

	pumpkins := make([]Pumpkin, n)
	pumpkins[0] = Pumpkin{
		Length:  0,
		Degrees: -1,
		Index:   1,
	}
	for i := 1; i < n; i++ {
		pumpkins[i].Create(fx, fy, i, in)
	}

	eps := 1e-10
	sort.Slice(pumpkins, func(i, j int) bool {
		d1 := pumpkins[i].Degrees
		d2 := pumpkins[j].Degrees
		if math.Abs(d1-d2) > eps {
			return d1 < d2
		}
		return pumpkins[i].Length < pumpkins[j].Length
	})

	startPoint := 1
	for i := 1; i < n-1; i++ {
		if pumpkins[i+1].Degrees-pumpkins[i].Degrees > 179.999 {
			startPoint = i + 1
			break
		}
	}

	fmt.Fprintln(out, n)
	fmt.Fprintln(out, 1)

	for i := startPoint; i < n; i++ {
		fmt.Fprintln(out, pumpkins[i].Index)
	}

	for i := 1; i < startPoint; i++ {
		fmt.Fprintln(out, pumpkins[i].Index)
	}
}
