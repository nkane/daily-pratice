package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func insertLevelOrder(a []interface{}, root *TreeNode, i int, n int) *TreeNode {
	if i < n {
		var temp *TreeNode
		if a[i] != nil {
			temp = &TreeNode{
				Val: a[i].(int),
			}
		}
		root = temp
		if root != nil {
			root.Left = insertLevelOrder(a, root.Left, 2*i+1, n)
			root.Right = insertLevelOrder(a, root.Right, 2*i+2, n)
		}
	}
	return root
}

func TestLowestCommonAncestor(t *testing.T) {
	var root *TreeNode
	a := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root = insertLevelOrder(a, root, 0, len(a))
	got := lowestCommonAncestor(root, root.Left, root.Right)
	assert.DeepEqual(t, root, got)

	a = []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root = insertLevelOrder(a, root, 0, len(a))
	got = lowestCommonAncestor(root, root.Left, root.Left.Right.Right)
	assert.DeepEqual(t, root.Left, got)
}

func TestLowestCommonAncestor_DFS(t *testing.T) {
	var root *TreeNode
	a := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root = insertLevelOrder(a, root, 0, len(a))
	got := lowestCommonAncestor_dfs(root, root.Left, root.Right)
	assert.DeepEqual(t, root, got)

	a = []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root = insertLevelOrder(a, root, 0, len(a))
	got = lowestCommonAncestor_dfs(root, root.Left, root.Left.Right.Right)
	assert.DeepEqual(t, root.Left, got)
}
