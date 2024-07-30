package leetcode

import "container/list"

/*
	104. Maximum Depth of Binary Tree
	Given the root of a binary tree, return it's maximum depth.
	A binary tree's maximum depth is the number of nodes along the longest path
	from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth_dfs_recursive(root *TreeNode) int {
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

func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			n, ok := q.Remove(q.Front()).(*TreeNode)
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

func maxDepth_dfs_preorder_iterative(root *TreeNode) int {
	stack := list.New()
	type Element struct {
		Node  *TreeNode
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
