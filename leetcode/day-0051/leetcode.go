package leetcode

import "math"

func minEatintgSpeed(piles []int, h int) int {
	left := 1
	right := 1
	for _, pile := range piles {
		right = Max(right, pile)
	}
	for left < right {
		mid := (left + right) / 2
		hourSpent := 0
		for _, pile := range piles {
			hourSpent += int(math.Ceil((float64(pile) / float64(mid))))
		}
		if hourSpent <= h {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return right
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
