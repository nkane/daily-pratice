package reverse_words

import "strings"

func reverseWords(s string) string {
	sb := strings.Builder{}
	words := strings.Fields(s)
	for i := len(words) - 1; i > 0; i-- {
		sb.WriteString(words[i] + " ")
	}
	sb.WriteString(words[0])
	return sb.String()
}

func reverseWordsFast(s string) string {
	arr := strings.Fields(s)
	lp, rp := 0, len(arr)-1
	for lp < rp {
		arr[lp], arr[rp] = arr[rp], arr[lp]
		lp++
		rp--
	}
	return strings.TrimSpace(strings.Join(arr, " "))
}
