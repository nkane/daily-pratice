package leetcode

/*
	374: Guess Number Higher or Lower
	API is provided by leetcode for guess(n int) int
*/

func guess(n int) int {
	secretNum := 6
	if n > secretNum {
		return -1
	} else if n < secretNum {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	hi := n
	lo := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g == -1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
