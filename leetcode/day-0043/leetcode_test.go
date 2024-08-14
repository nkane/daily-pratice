package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			Output: 4,
		},
		{
			Name: "Example 2",
			Grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			Output: -1,
		},
		{
			Name: "Example 3",
			Grid: [][]int{
				{0, 2},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := orangesRotting(test.Grid)
			assert.Assert(t, test.Output == got, "expected: %d, got: %d", test.Output, got)
		})
	}
}

func TestOrangesRottingInPlace(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			Output: 4,
		},
		{
			Name: "Example 2",
			Grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			Output: -1,
		},
		{
			Name: "Example 3",
			Grid: [][]int{
				{0, 2},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := orangesRottingInPlace(test.Grid)
			assert.Assert(t, test.Output == got, "expected: %d, got: %d", test.Output, got)
		})
	}
}

func TestOrangesRottingNeet(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			Output: 4,
		},
		{
			Name: "Example 2",
			Grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			Output: -1,
		},
		{
			Name: "Example 3",
			Grid: [][]int{
				{0, 2},
			},
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := orangesRottingNeet(test.Grid)
			assert.Assert(t, test.Output == got, "expected: %d, got: %d", test.Output, got)
		})
	}
}
