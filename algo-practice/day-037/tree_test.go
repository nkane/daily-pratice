package practice

import (
	"bytes"
	"io"
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

func OutputSetup(t *testing.T) (io.Reader, io.WriteCloser, *os.File) {
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w
	return r, w, oldOut
}

func OutputCleanup(t *testing.T, r io.Reader, w io.WriteCloser, oldOut *os.File) string {
	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	out := buf.String()
	return out
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		Expected string
	}{
		{
			Expected: "8 4 2 5 1 9 6 10 3 7 ",
		},
	}
	tree := CreateTestTree()
	r, w, oldOut := OutputSetup(t)
	for _, test := range tests {
		InorderTraversal(tree)
		got := OutputCleanup(t, r, w, oldOut)
		assert.Assert(t, test.Expected == got, "\nexpected: %s\ngot: %s", test.Expected, got)
	}
}

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		Expected string
	}{
		{
			Expected: "1 2 4 8 5 3 6 9 10 7 ",
		},
	}
	tree := CreateTestTree()
	for _, test := range tests {
		r, w, oldOut := OutputSetup(t)
		PreorderTraversal(tree)
		got := OutputCleanup(t, r, w, oldOut)
		assert.Assert(t, test.Expected == got, "\nexpected:\n%s, got:\n%s", test.Expected, got)
	}
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		Expected string
	}{
		{
			Expected: "8 4 5 2 9 10 6 7 3 1 ",
		},
	}
	tree := CreateTestTree()
	for _, test := range tests {
		r, w, oldOut := OutputSetup(t)
		PostorderTraversal(tree)
		got := OutputCleanup(t, r, w, oldOut)
		assert.Assert(t, test.Expected == got, "\nexpected:\n%s, got:\n%s", test.Expected, got)
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		Expected string
	}{
		{
			Expected: "1 2 3 4 5 6 7 8 9 10 ",
		},
	}
	tree := CreateTestTree()
	for _, test := range tests {
		r, w, oldOut := OutputSetup(t)
		BFS(tree)
		got := OutputCleanup(t, r, w, oldOut)
		assert.Assert(t, test.Expected == got, "expected:\n%s, got:\n%s", test.Expected, got)
	}
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

	tests := []struct {
		Expected int
	}{
		{
			Expected: 3,
		},
	}
	for _, test := range tests {
		got := MaxDepth_DFS_Recursive(&root)
		assert.Assert(t, got == test.Expected)

		got = MaxDepth_BFS(&root)
		assert.Assert(t, got == test.Expected)

		got = MaxDepth_DFS_Preorder_Iterative(&root)
		assert.Assert(t, got == test.Expected)
	}
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
