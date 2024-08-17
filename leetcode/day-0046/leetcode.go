package leetcode

import (
	"container/heap"
	"sort"
)

/*
	2542. Maximum Subsequence Score
*/

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	// Sort indices based on the descending order of nums2 values
	sort.Slice(indices, func(i, j int) bool {
		return nums2[indices[i]] > nums2[indices[j]]
	})

	h := &MinHeap{}
	heap.Init(h)
	sum := 0
	maxScore := int64(0)

	for i := 0; i < n; i++ {
		idx := indices[i]
		heap.Push(h, nums1[idx])
		sum += nums1[idx]

		if h.Len() > k {
			sum -= heap.Pop(h).(int)
		}

		if h.Len() == k {
			score := int64(sum) * int64(nums2[idx])
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
