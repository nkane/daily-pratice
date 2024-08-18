package leetcode

import "container/heap"

/*
	2462: Total Cost to Hire K Workers
*/

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinHeap) Peek() any {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

func totalCost(costs []int, k int, candidates int) int64 {
	headWorkers := &MinHeap{}
	tailWorkers := &MinHeap{}
	heap.Init(headWorkers)
	heap.Init(tailWorkers)
	for i := 0; i < candidates; i++ {
		heap.Push(headWorkers, costs[i])
	}
	for i := Max(candidates, len(costs)-candidates); i < len(costs); i++ {
		heap.Push(tailWorkers, costs[i])
	}
	answer := int64(0)
	nextHead := candidates
	nextTail := len(costs) - 1 - candidates
	for i := 0; i < k; i++ {
		if tailWorkers.Len() == 0 || (headWorkers.Len() > 0 && headWorkers.Peek().(int) <= tailWorkers.Peek().(int)) {
			answer += int64(heap.Pop(headWorkers).(int))
			// Only refill the queue if there are workers outside the two queues.
			if nextHead <= nextTail {
				heap.Push(headWorkers, costs[nextHead])
				nextHead++
			}
		} else {
			answer += int64(heap.Pop(tailWorkers).(int))
			// Only refill the queue if there are workers outside the two queues.
			if nextHead <= nextTail {
				heap.Push(tailWorkers, costs[nextTail])
				nextTail--
			}
		}
	}
	return answer
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
