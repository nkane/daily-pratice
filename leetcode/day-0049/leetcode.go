package leetcode

import (
	"sort"
)

/*
	2300: Successful Pairs of Spells and Potions
*/

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}
	n := len(potions)
	for _, spell := range spells {
		minPotion := (int(success) + spell - 1) / spell
		low, high := 0, n-1
		for low <= high {
			mid := low + (high-low)/2
			potion := potions[mid]
			if potion < minPotion {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		res = append(res, n-low)
	}
	return res
}
