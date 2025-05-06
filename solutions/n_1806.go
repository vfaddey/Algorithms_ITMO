package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	node int
	dist int64
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	T := make([]int64, 10)
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &T[i])
	}

	nums := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	src, dst := 0, n-1

	// Индексирование номеров
	idx := make(map[string]int, n)
	for i, s := range nums {
		idx[s] = i
	}

	const INF = int64(1e18)
	dist := make([]int64, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = INF
		prev[i] = -1
	}
	dist[src] = 0

	// Инициализация очереди
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: src, dist: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		u, d := item.node, item.dist
		if d != dist[u] {
			continue
		}
		if u == dst {
			break
		}
		s := nums[u]

		// Соседи: изменение одной цифры
		sBytes := []byte(s)
		for i := 0; i < 10; i++ {
			orig := sBytes[i]
			for c := byte('0'); c <= byte('9'); c++ {
				if c == orig {
					continue
				}
				sBytes[i] = c
				if v, ok := idx[string(sBytes)]; ok {
					newDist := d + T[i]
					if newDist < dist[v] {
						dist[v] = newDist
						prev[v] = u
						heap.Push(pq, Item{node: v, dist: newDist})
					}
				}
			}
			sBytes[i] = orig
		}

		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				if sBytes[i] == sBytes[j] {
					continue
				}
				sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
				if v, ok := idx[string(sBytes)]; ok {
					newDist := d + T[i]
					if newDist < dist[v] {
						dist[v] = newDist
						prev[v] = u
						heap.Push(pq, Item{node: v, dist: newDist})
					}
				}
				sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
			}
		}
	}

	if dist[dst] == INF {
		fmt.Println(-1)
		return
	}

	path := []int{}
	for cur := dst; cur != -1; cur = prev[cur] {
		path = append(path, cur+1)
	}
	reverse(path)

	fmt.Println(dist[dst])
	fmt.Println(len(path))
	for i, v := range path {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
