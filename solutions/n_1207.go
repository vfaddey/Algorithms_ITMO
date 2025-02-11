package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y float64
	idx  int
	ang  float64
}

func main() {

	var n int
	var x, y float64
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	points := make([]Point, n)
	var sumx, sumy float64
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		points[i] = Point{x: x, y: y, idx: i + 1}
		sumx += x
		sumy += y
	}
	cx := sumx / float64(n)
	cy := sumy / float64(n)
	for i := 0; i < n; i++ {
		points[i].ang = math.Atan2(points[i].y-cy, points[i].x-cx)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].ang < points[j].ang
	})

	halfCount := (n - 2) / 2
	eps := 1e-9

	for i := 0; i < n; i++ {
		j := (i + n/2) % n
		A := points[i]
		B := points[j]
		dx := B.x - A.x
		dy := B.y - A.y
		left := 0
		right := 0
		for k := 0; k < n; k++ {
			if k == i || k == j {
				continue
			}
			cp := dx*(points[k].y-A.y) - dy*(points[k].x-A.x)
			if cp > eps {
				left++
			} else if cp < -eps {
				right++
			}
		}
		if left == halfCount && right == halfCount {
			fmt.Println(A.idx, B.idx)
			return
		}
	}
}
