package leetcode

import (
	"container/heap"
	"fmt"
	"math/rand"
)

/*
	215: Kth Largest Element in an Array

	Given an integer array nums and an integer k, return the kth largest element in the array.

	Note that it is the kth largest element in sorted order, not the kth distinct element.

	Can you solve it without sorting?
*/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	// this makes it a min-heap
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

func minHeapify(heap []int, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < len(heap) && heap[left] < heap[smallest] {
		smallest = left
	}
	if right < len(heap) && heap[right] < heap[smallest] {
		smallest = right
	}
	if smallest != i {
		heap[i], heap[smallest] = heap[smallest], heap[i]
		minHeapify(heap, smallest)
	}
}

// Function to build a min-heap from a slice
func buildMinHeap(heap []int) {
	for i := len(heap)/2 - 1; i >= 0; i-- {
		minHeapify(heap, i)
	}
}

// Function to replace the root of the heap and maintain the min-heap property
func replaceMin(heap []int, val int) {
	heap[0] = val
	minHeapify(heap, 0)
}

func findKthLargest_Slice(nums []int, k int) int {
	// Initialize the min-heap with the first k elements
	heap := make([]int, k)
	copy(heap, nums[:k])

	// Build the initial min-heap
	buildMinHeap(heap)

	// Iterate over the rest of the elements
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			replaceMin(heap, nums[i])
		}
	}

	// The root of the heap is the kth largest element
	return heap[0]
}

// MinHeap represents a min-heap.
type MinHeap struct {
	heap []int
}

// NewMinHeap creates a new MinHeap instance.
func NewMinHeap() *MinHeap {
	return &MinHeap{heap: []int{}}
}

// Insert adds a new element to the heap.
func (h *MinHeap) Insert(val int) {
	h.heap = append(h.heap, val)
	h.heapifyUp(len(h.heap) - 1)
}

// ExtractMin removes and returns the smallest element from the heap.
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.heap) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	// The root of the heap is the minimum element.
	min := h.heap[0]

	// Move the last element to the root and heapify down.
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapifyDown(0)

	return min, nil
}

// heapifyUp ensures the heap property is maintained after insertion.
func (h *MinHeap) heapifyUp(index int) {
	for h.heap[parent(index)] > h.heap[index] {
		h.heap[parent(index)], h.heap[index] = h.heap[index], h.heap[parent(index)]
		index = parent(index)
	}
}

// heapifyDown ensures the heap property is maintained after extraction.
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.heap) - 1
	leftChild, rightChild := left(index), right(index)
	childToCompare := 0

	for leftChild <= lastIndex {
		if leftChild == lastIndex { // only left child exists
			childToCompare = leftChild
		} else if h.heap[leftChild] < h.heap[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		// If the current node is smaller than the smallest child, the heap is valid.
		if h.heap[index] <= h.heap[childToCompare] {
			return
		}

		// Swap and continue heapifying down.
		h.heap[index], h.heap[childToCompare] = h.heap[childToCompare], h.heap[index]
		index = childToCompare
		leftChild, rightChild = left(index), right(index)
	}
}

// parent returns the index of the parent of the given node.
func parent(index int) int {
	return (index - 1) / 2
}

// left returns the index of the left child of the given node.
func left(index int) int {
	return 2*index + 1
}

// right returns the index of the right child of the given node.
func right(index int) int {
	return 2*index + 2
}

// Peek returns the smallest element without removing it.
func (h *MinHeap) Peek() (int, error) {
	if len(h.heap) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.heap[0], nil
}

// Size returns the number of elements in the heap.
func (h *MinHeap) Size() int {
	return len(h.heap)
}

func findKthLargest_QuickSelect(nums []int, k int) int {
	return QuickSelect(nums, k)
}

func QuickSelect(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	pivot := nums[rand.Intn(len(nums))]
	left := []int{}
	mid := []int{}
	right := []int{}
	for _, num := range nums {
		if num > pivot {
			left = append(left, num)
		} else if num < pivot {
			right = append(right, num)
		} else {
			mid = append(mid, num)
		}
	}
	if k <= len(left) {
		return QuickSelect(left, k)
	}
	if len(left)+len(mid) < k {
		return QuickSelect(right, k-len(left)-len(mid))
	}
	return pivot
}
