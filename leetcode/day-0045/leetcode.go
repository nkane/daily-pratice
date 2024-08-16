package leetcode

import (
	"container/heap"
	"sort"
)

type SmallestInfiniteSet struct {
	currentInteger int
	isPresent      map[int]bool
	addedIntegers  *MinHeap
}

func Constructor() SmallestInfiniteSet {
	h := &MinHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		currentInteger: 1,
		isPresent:      make(map[int]bool),
		addedIntegers:  h,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	var answer int
	if s.addedIntegers.Len() > 0 {
		answer = heap.Pop(s.addedIntegers).(int)
		delete(s.isPresent, answer)
	} else {
		answer = s.currentInteger
		s.currentInteger++
	}
	return answer
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num >= s.currentInteger || s.isPresent[num] {
		return
	}
	heap.Push(s.addedIntegers, num)
	s.isPresent[num] = true
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet_Sort struct {
	addedIntegers  map[int]struct{}
	currentInteger int
}

func Constructor_Sort() SmallestInfiniteSet_Sort {
	return SmallestInfiniteSet_Sort{
		addedIntegers:  make(map[int]struct{}),
		currentInteger: 1,
	}
}

func (s *SmallestInfiniteSet_Sort) PopSmallest() int {
	var answer int
	if len(s.addedIntegers) > 0 {
		// Convert map keys to a slice and sort it to get the smallest element
		keys := make([]int, 0, len(s.addedIntegers))
		for key := range s.addedIntegers {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		// The smallest element is the first one in the sorted slice
		answer = keys[0]
		delete(s.addedIntegers, answer)
	} else {
		answer = s.currentInteger
		s.currentInteger++
	}
	return answer
}
