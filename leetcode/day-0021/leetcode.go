package leetcode

import "sort"

/*
	1657. Determine if Two Strings are Close
	Two strings are considered close if you can attain one from the other using the following operations:
	- Operation 1: swap any two existing characters
		- For example, abcde -> aecdb
	- Operation 2: transform every occurence of one existing character into another existing character,
	and do the same with the other character.
		- For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
	You can use the operations on either string as many times as necessary.

	Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
*/

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Map := map[rune]int{}
	for _, c := range word1 {
		word1Map[c]++
	}
	word2Map := map[rune]int{}
	for _, c := range word2 {
		word2Map[c]++
	}

	if !equalKeySets(word1Map, word2Map) {
		return false
	}

	word1FrequencyList := mapValuesToList(word1Map)
	word2FrequencyList := mapValuesToList(word2Map)
	sort.Ints(word1FrequencyList)
	sort.Ints(word2FrequencyList)
	return equalSlices(word1FrequencyList, word2FrequencyList)
}

func equalKeySets(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k := range map1 {
		if _, found := map2[k]; !found {
			return false
		}
	}

	return true
}

func mapValuesToList[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func equalSlices[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func closeStrings_bit(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Map := make([]int, 26)
	word2Map := make([]int, 26)
	word1Bit := 0
	word2Bit := 0
	for _, c := range word1 {
		word1Map[c-'a']++
		word1Bit |= 1 << (c - 'a')
	}
	for _, c := range word2 {
		word2Map[c-'a']++
		word2Bit |= 1 << (c - 'a')
	}
	if word1Bit != word2Bit {
		return false
	}
	sort.Ints(word1Map)
	sort.Ints(word2Map)
	for i := 0; i < 26; i++ {
		if word1Map[i] != word2Map[i] {
			return false
		}
	}
	return true
}
