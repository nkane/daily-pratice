package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGoodNodes(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	expected := 4
	got := goodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = goodNodes_dfs(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	expected = 3
	got = goodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = goodNodes_dfs(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode{
		Val: 1,
	}
	expected = 1
	got = goodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = goodNodes_dfs(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode{
		Val: 9,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	expected = 1
	got = goodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = goodNodes_dfs(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
}
