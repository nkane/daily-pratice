package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](list []T, target T) (int, int) {
	count := 0
	lo := 0
	hi := len(list) - 1
	for lo <= hi {
		count++
		mid := (lo + hi) / 2
		guess := list[mid]
		if guess == target {
			return mid, count
		}
		if guess > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return 1, count
}
