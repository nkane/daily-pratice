package practice

import (
	"container/list"
	"fmt"
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
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		n, ok := q.Remove(q.Front()).(*TreeNode[T])
		if ok {
			fmt.Print(n.Data, " ")
			if n.Left != nil {
				q.PushBack(n.Left)
			}
			if n.Right != nil {
				q.PushBack(n.Right)
			}
		}
	}
}

func maxDepth_dfs_recursive[T any](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth_dfs_recursive(root.Left)
	rightMax := maxDepth_dfs_recursive(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}

func maxDepth_bfs[T any](root *TreeNode[T]) int {
	if root == nil {
		return 0
	}
	level := 0
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			n, ok := q.Remove(q.Front()).(*TreeNode[T])
			if ok {
				if n.Left != nil {
					q.PushBack(n.Left)
				}
				if n.Right != nil {
					q.PushBack(n.Right)
				}
			}
		}
		level += 1
	}
	return level
}

func maxDepth_dfs_preorder_iterative[T any](root *TreeNode[T]) int {
	stack := list.New()
	type Element struct {
		Node  *TreeNode[T]
		Depth int
	}
	stack.PushFront(Element{
		Node:  root,
		Depth: 1,
	})
	result := 0
	for stack.Len() > 0 {
		element, ok := stack.Remove(stack.Back()).(Element)
		if ok {
			if element.Node != nil {
				if result < element.Depth {
					result = element.Depth
				}
				stack.PushFront(Element{
					Node:  element.Node.Left,
					Depth: element.Depth + 1,
				})
				stack.PushFront(Element{
					Node:  element.Node.Right,
					Depth: element.Depth + 1,
				})
			}
		}
	}
	return result
}
