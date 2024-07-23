package practice

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/assert"
)

var printGraphs = false

func TestInorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	os.Stdout = w

	if printGraphs {
		DrawSVG(tree, "inorder_traversal")
	}

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

	if printGraphs {
		DrawSVG(tree, "preorder_traversal")
	}

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

	if printGraphs {
		DrawSVG(tree, "postorder_traversal")
	}

	PostorderTraversal(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "8 4 5 2 9 10 6 7 3 1 "
	assert.Assert(t, expected == got, "\nexpected:\n%s, got:\n%s", expected, got)
}

func TestLevelOrderTraversal(t *testing.T) {
	tree := CreateTestTree()
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	assert.NilError(t, err)
	os.Stdout = w

	if printGraphs {
		DrawSVG(tree, "levelorder_traversal")
	}

	LevelorderTraversal(tree)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "1 2 3 4 5 6 7 8 9 10 "
	assert.Assert(t, expected == got, "expected:\n%s, got:\n%s", expected, got)
}
