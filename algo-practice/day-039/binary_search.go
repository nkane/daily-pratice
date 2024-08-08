package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](a []T, target T) (int, int) {
	count := 0
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		count++
		mid := (lo + hi) / 2
		guess := a[mid]
		if target == guess {
			return mid, count
		}
		if target > guess {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1, count
}
