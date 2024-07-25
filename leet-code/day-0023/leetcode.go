package leetcode

import "strings"

/*
	2390. Removing Stars From a String
	You are given a string s, which contains starts *.
	In one operation, you can:
	- choose a start in s.
	- remove the closest non-start character to its left, as well as remove the start itself
	Return the string after all starts have been removed

	Note:
	- the input will be generated such that the operation is always possible.
	- it can be shown that the resulting string will always be unique.
*/

func removeStars(s string) string {
	r := []rune{}
	for i, c := range s {
		if c == '*' {
			if i > 0 {
				// remove last letter
				r = r[:len(r)-1]
			}
		} else {
			r = append(r, c)
		}
	}
	return string(r)
}

func removeStars_two_pointers(s string) string {
	ch := make([]rune, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c == '*' {
			j--
		} else {
			ch[j] = c
			j++
		}
	}
	sb := strings.Builder{}
	for i := 0; i < j; i++ {
		sb.WriteRune(ch[i])
	}
	return sb.String()
}
