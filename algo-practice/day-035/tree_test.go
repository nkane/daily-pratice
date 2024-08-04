package practice

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

func CreateTestTree() *TreeNode[int] {
	root := &TreeNode[int]{
		Data: 1,
	}
	// left subtree
	root.Left = &TreeNode[int]{
		Data:   2,
		Parent: root,
	}
	root.Left.Left = &TreeNode[int]{
		Data:   4,
		Parent: root.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		Data:   8,
		Parent: root.Left.Left,
	}
	root.Left.Right = &TreeNode[int]{
		Data:   5,
		Parent: root.Left,
	}
	// right subtree
	root.Right = &TreeNode[int]{
		Data:   3,
		Parent: root,
	}
	root.Right.Left = &TreeNode[int]{
		Data:   6,
		Parent: root.Right,
	}
	root.Right.Left.Left = &TreeNode[int]{
		Data:   9,
		Parent: root.Right.Left,
	}
	root.Right.Left.Right = &TreeNode[int]{
		Data:   10,
		Parent: root.Right.Left,
	}
	root.Right.Right = &TreeNode[int]{
		Data:   7,
		Parent: root.Right,
	}
	return root
}

func TestInorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w

	InorderTraversal(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "8 4 2 5 1 9 6 10 3 7 "
	assert.Assert(t, got == expected, "\nexpected: %s\ngot: %s", expected, got)
}

func TestPreorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w

	PreorderTraversal(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "1 2 4 8 5 3 6 9 10 7 "
	assert.Assert(t, expected == got, "\nexpected:\n%s, got:\n%s", expected, got)
}

func TestPostorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	assert.NilError(t, err)
	os.Stdout = w

	PostorderTraversal(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "8 4 5 2 9 10 6 7 3 1 "
	assert.Assert(t, expected == got, "\nexpected:\n%s, got:\n%s", expected, got)
}

func TestBFS(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	assert.NilError(t, err)
	os.Stdout = w

	BFS(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "1 2 3 4 5 6 7 8 9 10 "
	assert.Assert(t, expected == got, "expected:\n%s, got:\n%s", expected, got)
}

func TestMaxDepth(t *testing.T) {
	root := TreeNode[int]{
		Data: 3,
		Left: &TreeNode[int]{
			Data: 9,
		},
		Right: &TreeNode[int]{
			Data: 20,
			Left: &TreeNode[int]{
				Data: 15,
			},
			Right: &TreeNode[int]{
				Data: 7,
			},
		},
	}

	expected := 3
	got := MaxDepth_DFS_Recursive(&root)
	assert.Assert(t, got == expected)

	got = MaxDepth_BFS(&root)
	assert.Assert(t, got == expected)

	got = MaxDepth_DFS_Preorder_Iterative(&root)
	assert.Assert(t, got == expected)
}

// 1448: Count Good Nodes in Binary Tree
// Given a binary tree `root`, a node X in the tree is named good if in the path from root to X
// there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.
func TestGoodNodes(t *testing.T) {
	root := &TreeNode[int]{
		Data: 3,
		Left: &TreeNode[int]{
			Data: 1,
			Left: &TreeNode[int]{
				Data: 3,
			},
		},
		Right: &TreeNode[int]{
			Data: 4,
			Left: &TreeNode[int]{
				Data: 1,
			},
			Right: &TreeNode[int]{
				Data: 5,
			},
		},
	}
	expected := 4
	got := GoodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = GoodNodes_DFS(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode[int]{
		Data: 3,
		Left: &TreeNode[int]{
			Data: 3,
			Left: &TreeNode[int]{
				Data: 4,
			},
			Right: &TreeNode[int]{
				Data: 2,
			},
		},
	}
	expected = 3
	got = GoodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = GoodNodes_DFS(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode[int]{
		Data: 1,
	}
	expected = 1
	got = GoodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = GoodNodes_DFS(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)

	root = &TreeNode[int]{
		Data: 9,
		Right: &TreeNode[int]{
			Data: 3,
			Left: &TreeNode[int]{
				Data: 6,
			},
		},
	}
	expected = 1
	got = GoodNodes(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
	got = GoodNodes_DFS(root)
	assert.Assert(t, expected == got, "expected %d, got %d\n", expected, got)
}
