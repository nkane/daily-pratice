package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](a []T, x T) (int, int) {
	lo := 0
	hi := len(a) - 1
	count := 0
	mid := 0
	for lo <= hi {
		count++
		mid = (lo + hi) / 2
		guess := a[mid]
		if guess == x {
			return mid, count
		}
		if guess > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1, count
}
