package practice

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/golang-collections/collections/queue"
)

type TreeNode[T any] struct {
	ID     int
	Data   T
	Right  *TreeNode[T]
	Left   *TreeNode[T]
	Parent *TreeNode[T]
}

func CreateTestTree() *TreeNode[int] {
	root := &TreeNode[int]{
		ID:   1,
		Data: 1,
	}
	// left subtree
	root.Left = &TreeNode[int]{
		ID:     2,
		Data:   2,
		Parent: root,
	}
	root.Left.Left = &TreeNode[int]{
		ID:     4,
		Data:   4,
		Parent: root.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		ID:     8,
		Data:   8,
		Parent: root.Left.Left,
	}
	root.Left.Right = &TreeNode[int]{
		ID:     5,
		Data:   5,
		Parent: root.Left,
	}
	// right subtree
	root.Right = &TreeNode[int]{
		ID:     3,
		Data:   3,
		Parent: root,
	}
	root.Right.Left = &TreeNode[int]{
		ID:     6,
		Data:   6,
		Parent: root.Right,
	}
	root.Right.Left.Left = &TreeNode[int]{
		ID:     9,
		Data:   9,
		Parent: root.Right.Left,
	}
	root.Right.Left.Right = &TreeNode[int]{
		ID:     10,
		Data:   10,
		Parent: root.Right.Left,
	}
	root.Right.Right = &TreeNode[int]{
		ID:     7,
		Data:   7,
		Parent: root.Right,
	}
	return root
}

func DrawSVG[T any](root *TreeNode[T], name string) {
	gv := graph.New(graph.IntHash, graph.Directed())
	q := queue.New()
	q.Enqueue(root)
	var n *TreeNode[T]
	for q.Len() > 0 {
		n = q.Dequeue().(*TreeNode[T])
		if n != nil {
			gv.AddVertex(n.ID)
			if n.Left != nil {
				gv.AddVertex(n.Left.ID)
				gv.AddEdge(n.ID, n.Left.ID)
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				gv.AddVertex(n.Right.ID)
				gv.AddEdge(n.ID, n.Right.ID)
				q.Enqueue(n.Right)
			}
		}
	}
	file, err := os.Create(fmt.Sprintf("tree-%s.gv", name))
	if err != nil {
		panic(err)
	}
	_ = draw.DOT(gv, file)
	cmd := exec.Command("/usr/bin/dot", "-Tsvg", "-Kdot", "-O", fmt.Sprintf("tree-%s.gv", name))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func InorderTraversal[T any](node *TreeNode[T]) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Print(node.Data, " ")
		InorderTraversal(node.Right)
	}
}

func PreorderTraversal[T any](node *TreeNode[T]) {
	if node != nil {
		fmt.Print(node.Data, " ")
		PreorderTraversal(node.Left)
		PreorderTraversal(node.Right)
	}
}

func PostorderTraversal[T any](node *TreeNode[T]) {
	if node != nil {
		PostorderTraversal(node.Left)
		PostorderTraversal(node.Right)
		fmt.Print(node.Data, " ")
	}
}

func BFS[T any](node *TreeNode[T]) {
	if node == nil {
		return
	}
	q := queue.New()
	q.Enqueue(node)
	for q.Len() > 0 {
		n, ok := q.Dequeue().(*TreeNode[T])
		if ok {
			fmt.Print(n.Data, " ")
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
	}
}
