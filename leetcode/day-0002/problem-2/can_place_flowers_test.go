package can_place_flowers

import (
	"testing"

	"gotest.tools/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	a := []int{1, 0, 0, 0, 1}
	got := canPlaceFlowers(a, 1)
	assert.Assert(t, true == got)

	got = canPlaceFlowers(a, 2)
	assert.Assert(t, false == got)

	a = []int{1, 0, 0, 0, 0, 1}
	got = canPlaceFlowers(a, 2)
	assert.Assert(t, false == got)

	a = []int{0, 0, 1, 0, 1}
	got = canPlaceFlowers(a, 1)
	assert.Assert(t, true == got)

	a = []int{1, 0, 0, 0, 1, 0, 0}
	got = canPlaceFlowers(a, 2)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowers(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0, 1, 0, 0}
	got = canPlaceFlowers(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0}
	got = canPlaceFlowers(a, 1)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowers(a, 0)
	assert.Assert(t, true == got)
}

func TestCanPlaceFlowersRevised(t *testing.T) {
	a := []int{1, 0, 0, 0, 1}
	got := canPlaceFlowersRevised(a, 1)
	assert.Assert(t, true == got)

	got = canPlaceFlowersRevised(a, 2)
	assert.Assert(t, false == got)

	a = []int{1, 0, 0, 0, 0, 1}
	got = canPlaceFlowersRevised(a, 2)
	assert.Assert(t, false == got)

	a = []int{0, 0, 1, 0, 1}
	got = canPlaceFlowersRevised(a, 1)
	assert.Assert(t, true == got)

	a = []int{1, 0, 0, 0, 1, 0, 0}
	got = canPlaceFlowersRevised(a, 2)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowersRevised(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0, 1, 0, 0}
	got = canPlaceFlowersRevised(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0}
	got = canPlaceFlowersRevised(a, 1)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowersRevised(a, 0)
	assert.Assert(t, true == got)
}

func TestCanPlaceFlowersRevisedOptimized(t *testing.T) {
	a := []int{1, 0, 0, 0, 1}
	got := canPlaceFlowersRevisedOptimized(a, 1)
	assert.Assert(t, true == got)

	got = canPlaceFlowersRevisedOptimized(a, 2)
	assert.Assert(t, false == got)

	a = []int{1, 0, 0, 0, 0, 1}
	got = canPlaceFlowersRevisedOptimized(a, 2)
	assert.Assert(t, false == got)

	a = []int{0, 0, 1, 0, 1}
	got = canPlaceFlowersRevisedOptimized(a, 1)
	assert.Assert(t, true == got)

	a = []int{1, 0, 0, 0, 1, 0, 0}
	got = canPlaceFlowersRevisedOptimized(a, 2)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowersRevisedOptimized(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0, 1, 0, 0}
	got = canPlaceFlowersRevisedOptimized(a, 1)
	assert.Assert(t, true == got)

	a = []int{0, 0}
	got = canPlaceFlowersRevisedOptimized(a, 1)
	assert.Assert(t, true == got)

	a = []int{0}
	got = canPlaceFlowersRevisedOptimized(a, 0)
	assert.Assert(t, true == got)
}
