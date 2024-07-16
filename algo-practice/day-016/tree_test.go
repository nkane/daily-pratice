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
