package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := &IntHeap{100500, 2, 3, 5, 20, 4}
	heap.Init(h)
	heap.Push(h, 1000)
	fmt.Printf("min %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d\n", h.Pop())
	}
}

// IntHeap type
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
