package practice

import (
	"bytes"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGraphDFS(t *testing.T) {
	g := Graph[int, int]{
		Visited:       make(map[int]struct{}),
		AdjacencyList: make(map[int][]int),
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	oldOut := os.Stdout
	r, w, err := os.Pipe()
	assert.NilError(t, err)
	assert.NilError(t, err)
	os.Stdout = w

	g.DFS(2)

	w.Close()
	os.Stdout = oldOut
	buf := bytes.Buffer{}
	buf.ReadFrom(r)
	got := buf.String()
	expected := "2 0 1 3 "
	assert.Assert(t, expected == got, "\nexpected:\n%s, got:\n%s", expected, got)
}
