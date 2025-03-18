package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	typeID int
	count  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].count > pq[j].count }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func solve(k int, counts []int) []int {
	pq := make(PriorityQueue, 0)
	for i := 0; i < k; i++ {
		if counts[i] > 0 {
			heap.Push(&pq, &Item{typeID: i + 1, count: counts[i]})
		}
	}

	result := make([]int, 0)
	var prevType int

	for pq.Len() > 0 {
		top := heap.Pop(&pq).(*Item)

		if pq.Len() == 0 {
			for i := 0; i < top.count; i++ {
				result = append(result, top.typeID)
			}
			break
		}

		if prevType == 0 || top.typeID != prevType {
			result = append(result, top.typeID)
			top.count--
			if top.count > 0 {
				heap.Push(&pq, top)
			}
			prevType = top.typeID
		} else {
			second := heap.Pop(&pq).(*Item)
			result = append(result, second.typeID)
			second.count--
			if second.count > 0 {
				heap.Push(&pq, second)
			}
			heap.Push(&pq, top)
			prevType = second.typeID
		}
	}

	return result
}

func main() {
	var k int
	fmt.Scan(&k)
	counts := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&counts[i])
	}

	sequence := solve(k, counts)
	for i, id := range sequence {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(id)
	}
	fmt.Println()
}
