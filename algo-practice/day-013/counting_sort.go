package practice

func CountingSort(a []int) []int {
	m := 0
	// find the max
	for i := 0; i < len(a); i++ {
		if m < a[i] {
			m = a[i]
		}
	}
	count := make([]int, m+1)
	for i := 0; i < len(a)-1; i++ {
		j := a[i]
		count[j]++
	}
	for i := 1; i <= m; i++ {
		count[i] = count[i] + count[i-1]
	}
	output := make([]int, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		j := a[i]
		output[count[j]] = a[i]
		count[j] = count[j] - 1
	}
	return output
}
