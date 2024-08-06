package practice

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T any] struct {
	Data   T
	Right  *TreeNode[T]
	Left   *TreeNode[T]
	Parent *TreeNode[T]
}

func InorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Print(root.Data, " ")
		InorderTraversal(root.Right)
	}
}

func PreorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		fmt.Print(root.Data, " ")
		PreorderTraversal(root.Left)
		PreorderTraversal(root.Right)
	}
}

func PostorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		PostorderTraversal(root.Left)
		PostorderTraversal(root.Right)
		fmt.Print(root.Data, " ")
	}
}

func BFS[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	q := []*TreeNode[T]{
		root,
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		fmt.Print(n.Data, " ")
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
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
	if node == nil {
		return depth
	}
	q := []*TreeNode[T]{
		node,
	}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		depth++
	}
	return depth
}

func MaxDepth_DFS_Preorder_Iterative[T any](node *TreeNode[T]) int {
	level := 0
	type StackFrame struct {
		Node  *TreeNode[T]
		Level int
	}
	stack := []StackFrame{
		{
			Node:  node,
			Level: 1,
		},
	}
	for len(stack) > 0 {
		sf := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if sf.Node != nil {
			if sf.Level > level {
				level = sf.Level
			}
			if sf.Node.Left != nil {
				stack = append(stack, StackFrame{
					Node:  sf.Node.Left,
					Level: sf.Level + 1,
				})
			}
			if sf.Node.Right != nil {
				stack = append(stack, StackFrame{
					Node:  sf.Node.Right,
					Level: sf.Level + 1,
				})
			}
		}
	}
	return level
}

func GoodNodes[T constraints.Ordered](root *TreeNode[T]) int {
	type Frame struct {
		Node *TreeNode[T]
		Max  T
	}
	q := []Frame{
		{
			Node: root,
			Max:  root.Data,
		},
	}
	goodNodes := 0
	for len(q) > 0 {
		frame := q[0]
		q = q[1:]
		if frame.Node != nil {
			if frame.Max <= frame.Node.Data {
				goodNodes++
			}
			if frame.Node.Right != nil {
				m := frame.Max
				if m < frame.Node.Right.Data {
					m = frame.Node.Right.Data
				}
				q = append(q, Frame{
					Node: frame.Node.Right,
					Max:  m,
				})
			}
			if frame.Node.Left != nil {
				m := frame.Max
				if m < frame.Node.Left.Data {
					m = frame.Node.Left.Data
				}
				q = append(q, Frame{
					Node: frame.Node.Left,
					Max:  m,
				})
			}
		}
	}
	return goodNodes
}

func GoodNodes_DFS[T constraints.Ordered](root *TreeNode[T]) int {
	return DFS(root, root.Data)
}

func DFS[T constraints.Ordered](node *TreeNode[T], maxValue T) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Data >= maxValue {
		res = 1
		maxValue = node.Data
	}
	res += DFS(node.Left, maxValue)
	res += DFS(node.Right, maxValue)
	return res
}
