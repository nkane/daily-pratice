package practice

import (
	"container/list"
	"fmt"

	"golang.org/x/exp/constraints"
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
	max := rightMax
	if leftMax > rightMax {
		max = leftMax
	}
	return max
}

func MaxDepth_BFS[T any](root *TreeNode[T]) int {
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
		level++
	}
	return level
}

func MaxDepth_DFS_Preorder_Iterative[T any](root *TreeNode[T]) int {
	level := 0
	if root == nil {
		return level
	}
	type Element struct {
		Node  *TreeNode[T]
		Depth int
	}
	s := list.New()
	s.PushFront(Element{
		Node:  root,
		Depth: 1,
	})
	for s.Len() > 0 {
		e, ok := s.Remove(s.Front()).(Element)
		if !ok {
			return -1
		}
		if e.Node != nil {
			if e.Depth > level {
				level = e.Depth
			}
			s.PushFront(Element{
				Node:  e.Node.Left,
				Depth: e.Depth + 1,
			})
			s.PushFront(Element{
				Node:  e.Node.Right,
				Depth: e.Depth + 1,
			})
		}
	}
	return level
}

func goodNodes[T constraints.Ordered](root *TreeNode[T]) int {
	type Element struct {
		Node *TreeNode[T]
		Max  T
	}
	q := list.New()
	q.PushBack(Element{
		Node: root,
		Max:  root.Data,
	})
	goodNodes := 0
	for q.Len() > 0 {
		element, _ := q.Remove(q.Front()).(Element)
		if element.Node != nil {
			if element.Max <= element.Node.Data {
				goodNodes++
			}
			if element.Node.Right != nil {
				e := Element{
					Node: element.Node.Right,
				}
				e.Max = element.Max
				if e.Max < element.Node.Right.Data {
					e.Max = element.Node.Right.Data
				}
				q.PushBack(e)
			}
			if element.Node.Left != nil {
				e := Element{
					Node: element.Node.Left,
				}
				e.Max = element.Max
				if e.Max < element.Node.Left.Data {
					e.Max = element.Node.Left.Data
				}
				q.PushBack(e)
			}
		}
	}
	return goodNodes
}

func goodNodes_dfs[T constraints.Ordered](root *TreeNode[T]) int {
	return dfs(root, root.Data)
}

func dfs[T constraints.Ordered](node *TreeNode[T], maxValue T) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Data >= maxValue {
		res = 1
		maxValue = node.Data
	}
	res += dfs(node.Left, maxValue)
	res += dfs(node.Right, maxValue)
	return res
}
