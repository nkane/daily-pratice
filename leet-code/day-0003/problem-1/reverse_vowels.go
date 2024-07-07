package practice

import "strings"

func reverseVowels(s string) string {
	// spaces don't count
	vowelMap := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}
	frontIdx := 0
	backIdx := len(s) - 1
	// mid := (len(s) - 1) / 2
	// for frontIdx <= mid && backIdx >= 0 {
	for frontIdx <= backIdx && backIdx >= 0 {
		front := runes[frontIdx]
		back := runes[backIdx]
		_, frontOk := vowelMap[front]
		_, backOk := vowelMap[back]
		if frontOk && backOk {
			runes[frontIdx], runes[backIdx] = runes[backIdx], runes[frontIdx]
			frontIdx++
			backIdx--
		}
		if !frontOk {
			frontIdx++
		}
		if !backOk {
			backIdx--
		}
	}
	sb := strings.Builder{}
	sb.WriteString(string(runes))
	return sb.String()
}
