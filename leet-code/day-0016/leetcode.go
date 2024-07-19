package leetcode

/*
	1732. Find the Highest Altitude

	There is a biker going on a road trip. The road trip consist of `n + 1` points at different
	altitudes. The biker starts his trip on point `0` with an altitude equal `0`.

	You are given an integer array `gain` of length `n` where `gain[i]` is the net gain in altitude
	between points `i` and `i + 1` for all (0 <= i <= n). Return the highest altitude of a point.
*/

func largestAltitude(gain []int) int {
	altitude := 0
	max := 0
	for _, a := range gain {
		altitude += a
		if max < altitude {
			max = altitude
		}
	}
	return max
}

/*
Solution: Approach Prefix Sum

We start from the altitude 0 and we have a list of N integers, where each integer represents the gain in altitude
at each step (it could be negative as well, which implies a fall in altitude) a biker takes. We need to return the
highest altitude of the biker in the complete journey, including the starting point at 0.

This can be solved by taking the maximum altitudes at each step in the journey. The altitude at a step can be
determined as the altitude at the previous step plus the gain at the current step. Hence, we wills tart form 0 and
keep adding the gain in altitude to it at each step, and after each addition, we will update the maximum altitude
we have seen so far.

If we observe closely, the altitude at a point is the sum of the gains on the left of it, which is nothing but the
prefix sum at this index. Therefore, we can find the prefix sum and return the maximum as the highest reached altitudel.

Algorithm
1. Initialize the variable `currentAltitude` to `0`; this is the current altitude of the biker.
2. Initialize the variable `highestPoint` to `currentAltitude`, as the highest altitude we have seen is `0`.
3. Iterate over the gain in altitude in the list `gain` and add the current gain `altitudeGain`  to the variable
`currrentAltitude`.
4. Update the variable `highestPoint` as necessary.
5. return `highestPoint`.
*/

func largestAltitude_Solution(gain []int) int {
	currentAltitude := 0
	highestPoint := currentAltitude
	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain
		if highestPoint < currentAltitude {
			highestPoint = currentAltitude
		}
	}
	return highestPoint
}
