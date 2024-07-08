package practice

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](list []int, target int) (int, int) {
	lo := 0
	hi := len(list) - 1
	count := 0
	for lo <= hi {
		count++
		mid := (hi + lo) / 2
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
	return -1, count
}
