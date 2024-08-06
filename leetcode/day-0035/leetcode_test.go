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

func TestRightSideView(t *testing.T) {
	tests := []struct {
		Root     *TreeNode
		Expected []int
	}{
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Expected: []int{1, 3, 4},
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			Expected: []int{1, 3, 6, 4},
		},
		{
			Root:     nil,
			Expected: []int{},
		},
	}
	for _, test := range tests {
		got := rightSideView(test.Root)
		assert.DeepEqual(t, test.Expected, got)
		got = rightSideView_dfs(test.Root)
		assert.DeepEqual(t, test.Expected, got)
	}
}
