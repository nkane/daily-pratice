package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	magMap := map[rune]int{}
	for _, c := range magazine {
		magMap[c]++
	}
	for _, c := range ransomNote {
		if v, ok := magMap[c]; !ok || v <= 0 {
			return false
		}
		magMap[c]--
	}
	return true
}
