package main

import "fmt"

type heapNode struct {
	Imgidx int64 // imgidx
}

// indexheap，使用小顶堆，组织缓存
type indexheap struct {
	nodes             []heapNode
	Imgidx2importance *map[int64]float64
}

// functions necessary to build indexheap
// sort interface
func (h indexheap) Len() int { return int(len(h.nodes)) } // warning: must be int rather than int64
func (h indexheap) Less(i, j int) bool {
	return (*h.Imgidx2importance)[h.nodes[i].Imgidx] < (*h.Imgidx2importance)[h.nodes[j].Imgidx]
}
func (h indexheap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// heap interface push and pop
func (h *indexheap) Push(x interface{}) { (*h).nodes = append((*h).nodes, x.(heapNode)) }
func (h *indexheap) Pop() interface{} {
	old := (*h).nodes
	n := len(old)
	x := old[n-1]
	(*h).nodes = old[0 : n-1]
	return x
}

func main() {
	fmt.Println("ok!")
}
