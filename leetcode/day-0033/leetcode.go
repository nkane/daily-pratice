package leetcode

/*
	236: Lowest Common Ancestor of a Binary Tree
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Node       *TreeNode
	IsAncestor bool
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	r := ancestorHelper(root, p, q)
	if r.IsAncestor {
		return r.Node
	}
	return nil
}

func ancestorHelper(root, p, q *TreeNode) *Result {
	if root == nil {
		return &Result{
			Node:       nil,
			IsAncestor: false,
		}
	}
	if root == p && root == q {
		return &Result{
			Node:       root,
			IsAncestor: true,
		}
	}
	rx := ancestorHelper(root.Left, p, q)
	if rx.IsAncestor {
		return rx
	}
	ry := ancestorHelper(root.Right, p, q)
	if ry.IsAncestor {
		return ry
	}
	if rx.Node != nil && ry.Node != nil {
		return &Result{
			Node:       root,
			IsAncestor: true,
		}
	} else if root == p || root == q {
		isAncestor := false
		if rx.Node != nil || ry.Node != nil {
			isAncestor = true
		}
		return &Result{
			Node:       root,
			IsAncestor: isAncestor,
		}
	} else {
		var node *TreeNode
		if rx.Node != nil {
			node = rx.Node
		} else if ry.Node != nil {
			node = ry.Node
		}
		return &Result{
			Node:       node,
			IsAncestor: false,
		}
	}
}

func Covers(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	return Covers(root.Left, p) || Covers(root.Right, p)
}

func lowestCommonAncestor_dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor_dfs(root.Left, p, q)
	r := lowestCommonAncestor_dfs(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	if r != nil {
		return r
	}
	return nil
}
