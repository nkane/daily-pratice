package leetcode

/*
	1456: Maxmimum Number of Vowels in a Substring of Given Length
	Given a string `s` and an integer `k`, return the maximum number of vowels letters in
	any substring of `s` with length `k`.
*/

func maxVowels(s string, k int) int {
	max := 0
	for i := 0; i < k; i++ {
		max += checkVowel(rune(s[i]))
	}
	if max == k {
		return max
	}
	count := max
	for i := k; i < len(s); i++ {
		lastRune := rune(s[i-k])
		currentRune := rune(s[i])
		count -= checkVowel(lastRune)
		count += checkVowel(currentRune)
		if count > max {
			max = count
		}
		if max == k {
			return max
		}
	}
	return max
}

func checkVowel(c rune) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	case 'A', 'E', 'I', 'O', 'U':
		return 1
	}
	return 0
}

func maxVowels_solutions(s string, k int) int {
	m := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	count := 0
	for i := 0; i < k; i++ {
		if _, ok := m[rune(s[i])]; ok {
			count++
		}
	}
	answer := count
	for i := k; i < len(s); i++ {
		if _, ok := m[rune(s[i])]; ok {
			count += 1
		}
		if _, ok := m[rune(s[i-k])]; ok {
			count -= 1
		}
		if answer < count {
			answer = count
		}
	}
	return answer
}
