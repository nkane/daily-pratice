package leetcode

/*
	700: Search in a Binary Search Tree
	You are given the root of a binary search tree (BST) and an integer val.
	Find the node in the BST that the node's value equals val and return the subtree rooted with
	that node, if such a node does not exist, return null
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	q := []*TreeNode{
		root,
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.Val == val {
			return n
		}
		if val > n.Val {
			if n.Right != nil {
				q = append(q, n.Right)
			}
		} else {
			if n.Left != nil {
				q = append(q, n.Left)
			}
		}
	}
	return nil
}
