package leetcode

func isSubsequence(s string, t string) bool {
	idx := 0
	for i := 0; i <= len(t)-1 && idx < len(s); i++ {
		if t[i] == s[idx] {
			idx++
		}
	}
	result := len(s) == idx
	return result
}
