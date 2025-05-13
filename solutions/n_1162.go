package main

import (
	"bufio"
	"fmt"
	"os"
)

const eps = 1e-9

type Edge struct {
	from, to  int
	rate, com float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M, S int
	var V float64

	fmt.Fscan(reader, &N, &M, &S, &V)

	edges := make([]Edge, 0, 2*M)
	for i := 0; i < M; i++ {
		var A, B int
		var R_ab, C_ab, R_ba, C_ba float64
		fmt.Fscan(reader, &A, &B, &R_ab, &C_ab, &R_ba, &C_ba)
		edges = append(edges, Edge{from: A, to: B, rate: R_ab, com: C_ab})
		edges = append(edges, Edge{from: B, to: A, rate: R_ba, com: C_ba})
	}

	dist := make([]float64, N+1)
	dist[S] = V

	for i := 0; i < N; i++ {
		updated := false
		for _, e := range edges {
			if dist[e.from] > e.com+eps {
				amt := (dist[e.from] - e.com) * e.rate
				if amt > dist[e.to]+eps {
					dist[e.to] = amt
					updated = true
					if i == N-1 {
						fmt.Println("YES")
						return
					}
				}
			}
		}
		if !updated {
			break
		}
	}

	if dist[S] > V+eps {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
