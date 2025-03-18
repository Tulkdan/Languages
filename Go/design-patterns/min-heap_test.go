package main

import (
	"reflect"
	"testing"
)

func Test_MinHeap(t *testing.T) {
	t.Run("Creating heap push and pop", func(t *testing.T) {
		arr := []int{10, 7, 11, 5, 4, 13}
		heap := MinHeap{}

		for _, v := range arr {
			heap.Push(v)
		}

		expected := []int{4, 5, 11, 10, 7, 13}
		if !reflect.DeepEqual(heap.heap, expected) {
			t.Errorf("Min heap insertion is wrong, got %+v, expected %+v", heap.heap, expected)
		}

		heap.Pop()
		expected = []int{5, 7, 11, 10, 13}
		if !reflect.DeepEqual(heap.heap, expected) {
			t.Errorf("Min heap pop is wrong, got %+v, expected %+v", heap.heap, expected)
		}
	})

	t.Run("Passing array to constructor should be heapified", func(t *testing.T) {
		arr := []int{10, 7, 11, 5, 4, 13}
		heap := NewHeap(arr)

		expected := []int{4, 5, 11, 10, 7, 13}
		if !reflect.DeepEqual(heap.heap, expected) {
			t.Errorf("Min heap creation is wrong, got %+v, expected %+v", heap.heap, expected)
		}
	})
}
