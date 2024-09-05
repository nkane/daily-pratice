package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		Name   string
		Head   *ListNode
		Output *ListNode
	}{
		{
			Name: "Example 1",
			Head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			Output: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, test := range tests {
		got := middleNode(test.Head)
		assert.Assert(t, got != nil)
		n := test.Output
		gotN := got
		for n != nil {
			assert.Assert(t, n.Val == gotN.Val, "expected: %d, got: %d", n.Val, gotN.Val)
			n = n.Next
			gotN = gotN.Next
		}
	}
}
