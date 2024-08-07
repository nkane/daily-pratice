package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		Root     *TreeNode
		Expected int
	}{
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: -8,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			Expected: 2,
		},
		{
			Root: &TreeNode{
				Val: 1,
			},
			Expected: 1,
		},
	}
	for _, test := range tests {
		got := maxLevelSum(test.Root)
		assert.Assert(t, test.Expected == got, "expected: %+v, got %+v", test.Expected, got)
	}
}
