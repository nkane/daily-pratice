package practice

import "log"

type TreeNode[T any] struct {
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Parent *TreeNode[T]
	Data   T
}

func CreateTestTree() *TreeNode[int] {
	root := &TreeNode[int]{
		Data: 1,
	}
	// left subtree
	root.Left = &TreeNode[int]{
		Data:   2,
		Parent: root,
	}
	root.Left.Left = &TreeNode[int]{
		Data:   4,
		Parent: root.Left,
	}
	root.Left.Left.Left = &TreeNode[int]{
		Data:   8,
		Parent: root.Left.Left,
	}
	root.Left.Right = &TreeNode[int]{
		Data:   5,
		Parent: root.Left,
	}
	// right subtree
	root.Right = &TreeNode[int]{
		Data:   3,
		Parent: root,
	}
	root.Right.Left = &TreeNode[int]{
		Data:   6,
		Parent: root.Right,
	}
	root.Right.Left.Left = &TreeNode[int]{
		Data:   9,
		Parent: root.Right.Left,
	}
	root.Right.Left.Right = &TreeNode[int]{
		Data:   10,
		Parent: root.Right.Left,
	}
	root.Right.Right = &TreeNode[int]{
		Data:   7,
		Parent: root.Right,
	}
	return root
}

func InOrderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		InOrderTraversal(root.Left)
		log.Printf("%+v ", root.Data)
		InOrderTraversal(root.Right)
	}
}

func PreorderTraversal[T any](root *TreeNode[T]) {
	if root != nil {
		log.Printf("%+v ", root.Data)
		PreorderTraversal(root.Left)
		PreorderTraversal(root.Right)
	}
}
