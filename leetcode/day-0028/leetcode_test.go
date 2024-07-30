package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxDepth(t *testing.T) {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	expected := 3
	got := maxDepth_dfs_recursive(&root)
	assert.Assert(t, got == expected)

	got = maxDepth_bfs(&root)
	assert.Assert(t, got == expected)

	got = maxDepth_dfs_preorder_iterative(&root)
	assert.Assert(t, got == expected)
}
