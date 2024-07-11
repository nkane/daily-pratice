package practice

func CountSort(a []int) []int {
	n := len(a)
	m := 0
	for i := 0; i < n; i++ {
		if m < a[i] {
			m = a[i]
		}
	}
	countArray := make([]int, m+1)
	for i := 0; i < n; i++ {
		countArray[a[i]]++
	}
	for i := 1; i <= m; i++ {
		countArray[i] += countArray[i-1]
	}
	outputArray := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		outputArray[countArray[a[i]]-1] = a[i]
		countArray[a[i]]--
	}
	return outputArray
}
