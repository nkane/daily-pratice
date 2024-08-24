package practice

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T any] struct {
	Data   T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
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
	q := []*TreeNode[T]{
		node,
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n != nil {
			fmt.Print(n.Data, " ")
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
}

func MaxDepth_DFS_Recursive[T any](node *TreeNode[T]) int {
	if node == nil {
		return 0
	}
	maxLeft := MaxDepth_DFS_Recursive(node.Left) + 1
	maxRight := MaxDepth_DFS_Recursive(node.Right) + 1
	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight
}

func MaxDepth_BFS[T any](node *TreeNode[T]) int {
	depth := 0
	q := []*TreeNode[T]{
		node,
	}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]
			if n != nil {
				if n.Left != nil {
					q = append(q, n.Left)
				}
				if n.Right != nil {
					q = append(q, n.Right)
				}
			}
		}
		depth++
	}
	return depth
}

func MaxDepth_DFS_Preorder_Iterative[T any](node *TreeNode[T]) int {
	type StackFrame struct {
		Node  *TreeNode[T]
		Depth int
	}
	s := []StackFrame{
		{
			Node:  node,
			Depth: 1,
		},
	}
	depth := 0
	for len(s) > 0 {
		sf := s[len(s)-1]
		s = s[:len(s)-1]
		if sf.Node != nil {
			if sf.Depth > depth {
				depth = sf.Depth
			}
			if sf.Node.Left != nil {
				s = append(s, StackFrame{
					Node:  sf.Node.Left,
					Depth: sf.Depth + 1,
				})
			}
			if sf.Node.Right != nil {
				s = append(s, StackFrame{
					Node:  sf.Node.Right,
					Depth: sf.Depth + 1,
				})
			}
		}
	}
	return depth
}

func GoodNodes[T constraints.Ordered](node *TreeNode[T]) int {
	type Frame struct {
		Node *TreeNode[T]
		Max  T
	}
	q := []Frame{
		{
			Node: node,
			Max:  node.Data,
		},
	}
	goodNodes := 0
	for len(q) > 0 {
		f := q[0]
		q = q[1:]
		if f.Node != nil {
			if f.Node.Data >= f.Max {
				goodNodes++
			}
			if f.Node.Left != nil {
				q = append(q, Frame{
					Node: f.Node.Left,
					Max:  Max(f.Node.Left.Data, f.Max),
				})
			}
			if f.Node.Right != nil {
				q = append(q, Frame{
					Node: f.Node.Right,
					Max:  Max(f.Node.Right.Data, f.Max),
				})
			}
		}
	}
	return goodNodes
}

func Max[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func GoodNodes_DFS[T constraints.Ordered](node *TreeNode[T]) int {
	return DFS(node, node.Data)
}

func DFS[T constraints.Ordered](node *TreeNode[T], maxValue T) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Data >= maxValue {
		maxValue = node.Data
		res++
	}
	res += DFS(node.Left, maxValue)
	res += DFS(node.Right, maxValue)
	return res
}
