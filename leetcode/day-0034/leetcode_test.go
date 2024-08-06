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

func TestSearchBST(t *testing.T) {
	var root *TreeNode
	a := []interface{}{4, 2, 7, 1, 3}
	root = insertLevelOrder(a, root, 0, len(a))
	expected := root.Left
	got := searchBST(root, 2)
	assert.DeepEqual(t, expected, got)
}
