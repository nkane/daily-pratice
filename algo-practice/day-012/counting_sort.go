package practice

func CountingSort(a []int) []int {
	n := len(a)
	maxNum := 0
	for i := 0; i < n; i++ {
		if maxNum < a[i] {
			maxNum = a[i]
		}
	}
	count := make([]int, maxNum+1)
	for i := 0; i < n; i++ {
		count[a[i]]++
	}
	for i := 1; i <= maxNum; i++ {
		count[i] += count[i-1]
	}
	output := make([]int, len(a))
	for i := n - 1; i >= 0; i-- {
		output[count[a[i]]-1] = a[i]
		count[a[i]]--
	}
	return output
}
