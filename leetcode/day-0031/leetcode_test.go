package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestPathSum(t *testing.T) {
	a := []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
	var root *TreeNode
	root = insertLevelOrder(a, root, 0, len(a))
	expected := 3
	got := pathSum(root, 8)
	assert.Assert(t, expected == got)
	got = pathSum_optimized(root, 8)
	assert.Assert(t, expected == got)
}
