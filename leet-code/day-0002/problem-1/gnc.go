package gnc

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := slices.Max(candies)
	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		result[i] = (extraCandies+candies[i] >= max)
	}
	return result
}
