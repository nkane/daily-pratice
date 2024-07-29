package practice

import "golang.org/x/exp/constraints"

func Quicksort[T constraints.Ordered](a []T, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := Partition(a, lo, hi)
	Quicksort(a, lo, p-1)
	Quicksort(a, p+1, hi)
}

func Partition[T constraints.Ordered](a []T, lo int, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j <= hi-1; j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			i = i + 1
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
