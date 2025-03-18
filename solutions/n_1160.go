package main

import (
	"fmt"
	"sort"
)

// Edge представляет ребро графа
type Edge struct {
	u, v, w int // u и v — номера концентраторов, w — длина кабеля
}

// DSU представляет структуру данных для непересекающихся множеств
type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	return &DSU{parent}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px != py {
		d.parent[px] = py
	}
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		edges[i] = Edge{u, v, w}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	dsu := NewDSU(N)
	mstEdges := make([]Edge, 0)
	maxWeight := 0

	for _, edge := range edges {
		if dsu.Find(edge.u) != dsu.Find(edge.v) {
			dsu.Union(edge.u, edge.v)
			mstEdges = append(mstEdges, edge)
			if edge.w > maxWeight {
				maxWeight = edge.w
			}
		}
	}

	fmt.Println(maxWeight)
	fmt.Println(len(mstEdges))
	for _, edge := range mstEdges {
		fmt.Println(edge.u, edge.v)
	}
}
