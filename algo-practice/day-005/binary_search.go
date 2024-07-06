package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](list []T, x T) (int, int) {
	hi := len(list) - 1
	lo := 0
	count := 0
	mid := 0
	for lo <= hi {
		count++
		mid = (hi + lo) / 2
		guess := list[mid]
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
