package practice

import "golang.org/x/exp/constraints"

func Quicksort[T constraints.Ordered](a []T, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := Partition(a, lo, hi)
	Quicksort(a, p+1, hi)
	Quicksort(a, lo, p-1)
}

func Partition[T constraints.Ordered](a []T, lo int, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j <= hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
