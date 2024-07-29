package can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	availableSlots := 0
	if len(flowerbed) >= 3 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			availableSlots++
		}
		if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
			flowerbed[len(flowerbed)-1] = 1
			availableSlots++
		}
		// we're looking for 3 consecutive zeros
		for i := 0; i < len(flowerbed); i++ {
			if availableSlots >= n {
				return true
			}
			// look for 3 consecutive zeros
			limit := i + 3
			if limit > len(flowerbed)-1 {
				return availableSlots == n
			}
			if flowerbed[i] == 0 &&
				flowerbed[i+1] == 0 &&
				flowerbed[i+2] == 0 {
				flowerbed[i+1] = 1
				availableSlots++
			}
		}
	} else {
		if len(flowerbed) == 1 {
			if flowerbed[0] == 0 {
				availableSlots++
			}
		} else {
			if flowerbed[0] == 0 && flowerbed[1] == 0 {
				availableSlots++
			}
		}
	}
	return availableSlots == n
}

/*
The solution is very simple, we can find out the extra maxmium number of flowers, count, that can be
planed for the given flowerbed arragment. To do so, we can traverse over all the lements of the flowerbed
and find out those elements which are 0, implying an empty position. For every such element, we check if
its both adjacent positions are also empty. If so, we can plant a flower at the current position without
violating the no-adjacvent-flower-rule. For the first and last elements, we need not check the previous
and the next adjacent positions respectively.

If the count obtained is greater than or equal to n, the required number of flowers to be planted, we can
plant n flower in the empty spaces, otherwise not.$a

Time Complexity: O(n)
Space Complexity: O(1)
*/
func canPlaceFlowersRevised(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		// check if the current plot is empty
		if flowerbed[i] == 0 {
			// check if the left and right plots are empty
			emptyLeftPlot := (i == 0) || (flowerbed[i-1] == 0)
			emptyRightPlot := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
			// if both plots are empty, we can plant a flower here
			if emptyLeftPlot && emptyRightPlot {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}

/*
Instead of finding the maximum value of count that can be obtained, as done in the last approach, we can stop
the process of checking the positions for planting the flowers as soon as the count becomes equal to n. Doing this
leads to an optimization of the first approach. If count never becomes equal to n, n flowers can't be planted at
the empty positions.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func canPlaceFlowersRevisedOptimized(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		// check if the current plot is empty
		if flowerbed[i] == 0 {
			// check if the left and right plots are empty
			emptyLeftPlot := (i == 0) || (flowerbed[i-1] == 0)
			emptyRightPlot := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
			// if both plots are empty, we can plant a flower here
			if emptyLeftPlot && emptyRightPlot {
				flowerbed[i] = 1
				count++
				if count >= n {
					return true
				}
			}
		}
	}
	return count >= n
}
