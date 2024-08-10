package practice

import "golang.org/x/exp/constraints"

type TreeNode[T any] struct {
	Data   T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
}

func InorderTraversal[T any](node *TreeNode[T]) {
}

func PreorderTraversal[T any](node *TreeNode[T]) {
}

func PostorderTraversal[T any](node *TreeNode[T]) {
}

func BFS[T any](node *TreeNode[T]) {
}

func MaxDepth_DFS_Recursive[T any](node *TreeNode[T]) int {
	return -1
}

func MaxDepth_BFS[T any](node *TreeNode[T]) int {
	return -1
}

func MaxDepth_DFS_Preorder_Iterative[T any](node *TreeNode[T]) int {
	return -1
}

func GoodNodes[T constraints.Ordered](node *TreeNode[T]) int {
	return -1
}

func GoodNodes_DFS[T constraints.Ordered](node *TreeNode[T]) int {
	return -1
}
