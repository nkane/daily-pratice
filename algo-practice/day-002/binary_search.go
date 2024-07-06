package practice

func BinarySearch(list []int, x int) (int, int) {
	lo := 0
	hi := len(list) - 1
	count := 0
	for lo <= hi {
		count++
		mid := (lo + hi) / 2
		guess := list[mid]
		if guess == x {
			return mid, count
		}
		if guess > x {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1, count
}
