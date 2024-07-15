package practice

func CountingSort(a []int) []int {
	m := 0
	for _, e := range a {
		if m < e {
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
	for i := len(a) - 1; i >= 0; i-- {
		e := a[i]
		v := count[e] - 1
		count[e]--
		output[v] = e
	}
	return output
}
