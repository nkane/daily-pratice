package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		Name     string
		Rooms    [][]int
		Expected bool
	}{
		{
			Name: "Can Visit All Rooms",
			Rooms: [][]int{
				{
					1,
				},
				{
					2,
				},
				{
					3,
				},
				{},
			},
			Expected: true,
		},
		{
			Name: "Cannot Visit All Rooms",
			Rooms: [][]int{
				{
					1,
					3,
				},
				{
					3,
					0,
					1,
				},
				{
					2,
				},
				{
					0,
				},
			},
			Expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := canVisitAllRooms(test.Rooms)
			assert.Assert(t, test.Expected == got)
		})
	}
}
