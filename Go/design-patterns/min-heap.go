package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MinHeap struct {
	heap []int
}

func NewHeap(arr []int) *MinHeap {
	h := &MinHeap{heap: arr}
	for i := (len(arr) - 1)/2; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (c *MinHeap) Push(v int) {
	c.heap = append(c.heap, v)

	currIdx := len(c.heap) - 1

	for currIdx > 0 && c.heap[(currIdx - 1) / 2] > c.heap[currIdx] {
		c.swap((currIdx - 1) / 2, currIdx)
		currIdx = (currIdx - 1) / 2
	}
}

// Removes the min value (root)
func (c *MinHeap) Pop() int {
	if len(c.heap) == 0 {
		return -1
	}

	idx := 0
	root := c.heap[idx]
	c.heap[idx] = c.heap[len(c.heap) - 1]
	c.heap = c.heap[0:len(c.heap) - 1]

	c.heapify(idx)

	return root
}

func (c *MinHeap) heapify(idx int) {
	for true {
		left := 2 * idx + 1
		right := 2 * idx + 2
		smallest := idx

		if left < len(c.heap) && c.heap[left] < c.heap[smallest] {
			smallest = left
		}

		if right < len(c.heap) && c.heap[right] < c.heap[smallest] {
			smallest = right
		}

		if smallest != idx {
			c.swap(smallest, idx)
			idx = smallest
		} else {
			break
		}
	}
}

func (c *MinHeap) swap(from, to int) {
	c.heap[from], c.heap[to] = c.heap[to], c.heap[from]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		x, y := heap.Pop(h), heap.Pop(h)

		if x != y {
			heap.Push(h, x.(int) - y.(int))
		}
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

type Item struct {
	value    []int
	priority float64
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	h := &PriorityQueue{}
	heap.Init(h)

	for _, point := range points {
		distance := math.Sqrt(math.Pow(float64(point[0]), 2) + math.Pow(float64(point[1]), 2))
		item := &Item{
			value: point,
			priority: distance,
		}
		heap.Push(h, item)
	}

	result := [][]int{}

	for i := k; i > 0; i-- {
		info := heap.Pop(h).(Item)
		result = append(result, info.value)
	}

	return result
}



func minDaysToDeliverParcels(parcels []int32) int32 {
	var minToReduce int32 = 0

	for _, p := range parcels {
		if p > 0 {
			if minToReduce == 0 {
				minToReduce = p
			} else {
				minToReduce = min(minToReduce, p)
			}
		}
	}

	if minToReduce == 0 {
		return 0
	}

	for i, p := range parcels {
		if p > 0 {
			parcels[i] -= minToReduce
		}
	}

	return 1 + minDaysToDeliverParcels(parcels)
}

func main() {
	fmt.Println(minDaysToDeliverParcels([]int32{4, 2, 3, 4}))
	fmt.Println(minDaysToDeliverParcels([]int32{2,3,4,3,3}))
}



