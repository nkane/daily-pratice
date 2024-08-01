package leetcode

import "container/list"

/*
	1448: Count Good Nodes in Binary Tree
	Given a binary tree `root`, a node X in the tree is named good if in the path from root to X
	there are no nodes with a value greater than X.
	Return the number of good nodes in the binary tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	type Element struct {
		Node *TreeNode
		Max  int
	}
	q := list.New()
	q.PushBack(Element{
		Node: root,
		Max:  root.Val,
	})
	goodNodes := 0
	for q.Len() > 0 {
		element, _ := q.Remove(q.Front()).(Element)
		if element.Node != nil {
			if element.Max <= element.Node.Val {
				goodNodes++
			}
			if element.Node.Right != nil {
				e := Element{
					Node: element.Node.Right,
				}
				e.Max = element.Max
				if e.Max < element.Node.Right.Val {
					e.Max = element.Node.Right.Val
				}
				q.PushBack(e)
			}
			if element.Node.Left != nil {
				e := Element{
					Node: element.Node.Left,
				}
				e.Max = element.Max
				if e.Max < element.Node.Left.Val {
					e.Max = element.Node.Left.Val
				}
				q.PushBack(e)
			}
		}
	}
	return goodNodes
}

func goodNodes_dfs(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxValue int) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Val >= maxValue {
		res = 1
		maxValue = node.Val
	}
	res += dfs(node.Left, maxValue)
	res += dfs(node.Right, maxValue)
	return res
}
