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
	if node == nil {
		return
	}
	q := []*TreeNode[T]{
		node,
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
	depth := 0
	type StackFrame struct {
		Node  *TreeNode[T]
		Depth int
	}
	stack := []StackFrame{
		{
			Node:  node,
			Depth: 1,
		},
	}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.Depth > depth {
			depth = n.Depth
		}
		if n.Node.Left != nil {
			stack = append(stack, StackFrame{
				Node:  n.Node.Left,
				Depth: n.Depth + 1,
			})
		}
		if n.Node.Right != nil {
			stack = append(stack, StackFrame{
				Node:  n.Node.Right,
				Depth: n.Depth + 1,
			})
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
		frame := q[0]
		q = q[1:]
		if frame.Node != nil {
			if frame.Max <= frame.Node.Data {
				goodNodes++
			}
			if frame.Node.Right != nil {
				max := frame.Max
				if max < frame.Node.Right.Data {
					max = frame.Node.Right.Data
				}
				q = append(q, Frame{
					Node: frame.Node.Right,
					Max:  max,
				})
			}
			if frame.Node.Left != nil {
				max := frame.Max
				if max < frame.Node.Left.Data {
					max = frame.Node.Left.Data
				}
				q = append(q, Frame{
					Node: frame.Node.Left,
					Max:  max,
				})
			}
		}
	}
	return goodNodes
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
		res = 1
		maxValue = node.Data
	}
	res += DFS(node.Left, maxValue)
	res += DFS(node.Right, maxValue)
	return res
}
