package leetcode

/*
	450: Delete Node in BST
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val = cur.Val
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}

func deleteNode_iterative(root *TreeNode, key int) *TreeNode {
	node, parent := findNode(root, key)

	if node == nil {
		return root
	}

	if node == parent {
		if node.Left != nil {
			insert(node.Left, node.Right)
			return node.Left
		} else {
			insert(node.Right, node.Left)
			return node.Right
		}
	}

	isLeft := parent.Left == node

	if isLeft {
		if node.Left != nil {
			parent.Left = node.Left
			insert(node.Left, node.Right)
		} else {
			parent.Left = node.Right
		}
	} else {
		if node.Right != nil {
			parent.Right = node.Right
			insert(node.Right, node.Left)
		} else {
			parent.Right = node.Left
		}
	}

	return root
}

func insert(root *TreeNode, target *TreeNode) {
	if target == nil || root == nil {
		return
	}

	if target.Val > root.Val {
		if root.Right == nil {
			root.Right = target
		} else {
			insert(root.Right, target)
		}
	} else {
		if root.Left == nil {
			root.Left = target
		} else {
			insert(root.Left, target)
		}
	}
}

func findNode(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	parent := root
	for {
		if root == nil {
			return nil, nil
		}

		if root.Val == key {
			return root, parent
		}

		parent = root
		if root.Val > key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}
