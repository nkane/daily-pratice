package practice

func CountingSort(a []int) []int {
	m := 0
	for _, e := range a {
		if e > m {
			m = e
		}
	}
	count := make([]int, m+1)
	for _, e := range a {
		count[e]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	output := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		j := a[i]
		count[j]--
		output[count[j]] = a[i]
	}
	return output
}
