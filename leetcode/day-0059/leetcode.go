package leetcode

import "strconv"

/*
	412: Fizz Buzz
*/

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		idx := i - 1
		if (i % 3) == 0 {
			result[idx] += "Fizz"
		}
		if (i % 5) == 0 {
			result[idx] += "Buzz"
		}
		if result[idx] == "" {
			result[idx] = strconv.Itoa(i)
		}
	}
	return result
}

/*
	1342: Number of Steps to Reduce to a Number to Zero
*/

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
