package practice

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

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
