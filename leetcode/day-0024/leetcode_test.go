package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestAsteroidCollision(t *testing.T) {
	asteroids := []int{5, 10, -5}
	expected := []int{5, 10}
	got := asteroidCollision(asteroids)
	assert.DeepEqual(t, expected, got)

	asteroids = []int{8, -8}
	expected = []int{}
	got = asteroidCollision(asteroids)
	assert.DeepEqual(t, expected, got)

	asteroids = []int{10, 2, -5}
	expected = []int{10}
	got = asteroidCollision(asteroids)
	assert.DeepEqual(t, expected, got)

	asteroids = []int{-2, -1, 1, 2}
	expected = []int{-2, -1, 1, 2}
	got = asteroidCollision(asteroids)
	assert.DeepEqual(t, expected, got)

	asteroids = []int{-2, -2, 1, -2}
	expected = []int{-2, -2, -2}
	got = asteroidCollision(asteroids)
	assert.DeepEqual(t, expected, got)
}
