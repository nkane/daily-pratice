package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestLongestZigZag(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 6,
				},
				Left: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	expected := 3
	got := longestZigZag(root)
	assert.Assert(t, expected == got)

	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	expected = 4
	got = longestZigZag(root)
	assert.Assert(t, expected == got)
}
