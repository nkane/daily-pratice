package leetcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		key      int
		expected *TreeNode
	}{
		{
			name:     "delete leaf node",
			root:     &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}},
			key:      3,
			expected: &TreeNode{5, nil, &TreeNode{7, nil, nil}},
		},
		{
			name:     "delete node with one child",
			root:     &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{7, &TreeNode{6, nil, nil}, nil}},
			key:      7,
			expected: &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}},
		},
		{
			name: "delete node with two children",
			root: &TreeNode{
				5,
				&TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
				&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}},
			},
			key: 3,
			expected: &TreeNode{
				5,
				&TreeNode{4, &TreeNode{2, nil, nil}, nil},
				&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}},
			},
		},
		{
			name: "delete root node",
			root: &TreeNode{
				5,
				&TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
				&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}},
			},
			key: 5,
			expected: &TreeNode{
				6,
				&TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
				&TreeNode{7, nil, &TreeNode{8, nil, nil}},
			},
		},
		{
			name:     "delete node not in tree",
			root:     &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}},
			key:      10,
			expected: &TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deleteNode(tt.root, tt.key)
			assert.DeepEqual(t, result, tt.expected)
		})
	}
}
