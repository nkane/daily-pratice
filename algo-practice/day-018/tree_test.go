package practice

import (
	"bytes"
	"log"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestInorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetFlags(oldFlags)
		log.SetOutput(os.Stderr)
	}()
	InOrderTraversal(tree)
	got := buf.String()
	assert.Assert(t, got == "8 \n4 \n2 \n5 \n1 \n9 \n6 \n10 \n3 \n7 \n")
}

func TestPreorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetFlags(oldFlags)
		log.SetOutput(os.Stderr)
	}()
	PreorderTraversal(tree)
	got := buf.String()
	expected := "1 \n2 \n4 \n8 \n5 \n3 \n6 \n9 \n10 \n7 \n"
	assert.Assert(t, expected == got, "expected:\n%s, got:\n%s", expected, got)
}

func TestPostorderTraversal(t *testing.T) {
	tree := CreateTestTree()
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetFlags(oldFlags)
		log.SetOutput(os.Stderr)
	}()
	PostorderTraversal(tree)
	got := buf.String()
	expected := "8 \n4 \n5 \n2 \n9 \n10 \n6 \n7 \n3 \n1 \n"
	assert.Assert(t, expected == got, "expected:\n%s, got:\n%s", expected, got)
}

func TestLevelOrderTraversal(t *testing.T) {
	tree := CreateTestTree()
	buf := bytes.Buffer{}
	log.SetOutput(&buf)
	oldFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetFlags(oldFlags)
		log.SetOutput(os.Stderr)
	}()
	LevelorderTraversal(tree)
	got := buf.String()
	expected := "1 \n2 \n3 \n4 \n5 \n6 \n7 \n8 \n9 \n10 \n"
	assert.Assert(t, expected == got, "expected:\n%s, got:\n%s", expected, got)
}
