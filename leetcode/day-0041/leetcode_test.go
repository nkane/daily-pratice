package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		Name      string
		Equations [][]string
		Values    []float64
		Queries   [][]string
		Output    []float64
	}{
		{
			Name: "Example 1",
			Equations: [][]string{
				{"a", "b"},
				{"b", "c"},
			},
			Values: []float64{
				2.0,
				3.0,
			},
			Queries: [][]string{
				{"a", "c"},
				{"b", "a"},
				{"a", "e"},
				{"a", "a"},
				{"x", "x"},
			},
			Output: []float64{
				6.00000,
				0.50000,
				-1.00000,
				1.00000,
				-1.00000,
			},
		},
		{
			Name: "Example 2",
			Equations: [][]string{
				{"a", "b"},
				{"b", "c"},
				{"bc", "cd"},
			},
			Values: []float64{
				1.5,
				2.5,
				5.0,
			},
			Queries: [][]string{
				{"a", "c"},
				{"c", "b"},
				{"bc", "cd"},
				{"cd", "bc"},
			},
			Output: []float64{
				3.75000,
				0.40000,
				5.00000,
				0.20000,
			},
		},
		{
			Name: "Example 3",
			Equations: [][]string{
				{"a", "b"},
			},
			Values: []float64{
				0.5,
			},
			Queries: [][]string{
				{"a", "b"},
				{"b", "a"},
				{"a", "c"},
				{"x", "y"},
			},
			Output: []float64{
				0.50000,
				2.00000,
				-1.00000,
				-1.00000,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := calcEquation(test.Equations, test.Values, test.Queries)
			assert.DeepEqual(t, test.Output, got)
		})
	}
}
