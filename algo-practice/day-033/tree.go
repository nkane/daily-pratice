package practice

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T any] struct {
	ID     int
	Data   T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
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

func MaxDepth_DFS_Recursive[T any](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	leftMax := MaxDepth_DFS_Recursive(root.Left) + 1
	rightMax := MaxDepth_DFS_Recursive(root.Right) + 1
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

func MaxDepth_BFS[T any](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode[T]{
		root,
	}
	depth := 0
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

func MaxDepth_DFS_Preorder_Iterative[T any](root *TreeNode[T]) int {
	level := 0
	type Element struct {
		Node  *TreeNode[T]
		Level int
	}
	stack := []Element{
		{
			Node:  root,
			Level: 1,
		},
	}
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if e.Node != nil {
			if e.Node.Left != nil {
				stack = append(stack, Element{
					Node:  e.Node.Left,
					Level: e.Level + 1,
				})
			}
			if e.Node.Right != nil {
				stack = append(stack, Element{
					Node:  e.Node.Right,
					Level: e.Level + 1,
				})
			}
			if e.Level > level {
				level = e.Level
			}
		}
	}
	return level
}

func GoodNodes[T constraints.Ordered](root *TreeNode[T]) int {
	type Element struct {
		Node *TreeNode[T]
		Max  T
	}
	q := []Element{
		{
			Node: root,
			Max:  root.Data,
		},
	}
	goodNodes := 0
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		if e.Node != nil {
			if e.Max <= e.Node.Data {
				goodNodes++
			}
			newElement := Element{}
			if e.Node.Right != nil {
				newElement.Node = e.Node.Right
				newElement.Max = e.Max
				if newElement.Max < e.Node.Right.Data {
					newElement.Max = e.Node.Right.Data
				}
				q = append(q, newElement)
			}
			if e.Node.Left != nil {
				newElement.Node = e.Node.Left
				newElement.Max = e.Max
				if newElement.Max < e.Node.Left.Data {
					newElement.Max = e.Node.Left.Data
				}
				q = append(q, newElement)
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
