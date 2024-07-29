package leetcode

import "math"

/*
	725. Asteroid Collision
	We are given an array asteroids of integers representing asteroids in a row.

	For each astroid, the absolute value represents its size, and the sign represents
	its direction (positive meaning right, negative meaning left). Each asteroid moves
	at the same speed.

	Find out the state of the asteroids after all collisions. If two asteroids meet, the
	smaller one will explode. If both are the same size, both will explode. Two asteroids
	moving the same direction will never meet.
*/

func asteroidCollision(asteroids []int) []int {
	st := []int{}
	for _, asteroid := range asteroids {
		flag := true
		for len(st) > 0 && st[len(st)-1] > 0 && asteroid < 0 {
			if math.Abs(float64(st[len(st)-1])) < math.Abs(float64(asteroid)) {
				st = st[:len(st)-1]
				continue
			} else if math.Abs(float64(st[len(st)-1])) == math.Abs(float64(asteroid)) {
				st = st[:len(st)-1]
			}
			flag = false
			break
		}
		if flag {
			st = append(st, asteroid)
		}
	}
	return st
}
