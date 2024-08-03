package leetcode

/*
	1372: Longest ZigZag Path in a Binary Tree

	You are given a root of a binary tree.
	A ZigZig path for a binary tree is defined as follow:
	- Choose any node in the binary tree and a direction (right or left)
	- If the current direction is right, move to the right child or the current node; otherwise, move to the left
	child.
	- Change the direction from right to left or from left to right.
	- Repeat the second and third steps until you can't move in the tree.

	ZigZag length is defined as the number of nodes visisted -1, a single node has a length of 0.
	Return the longest ZigZag path contained in that tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	MoveLeft = iota
	MoveRight
)

func longestZigZag(root *TreeNode) int {
	max := 0
	type StackFrame struct {
		Node   *TreeNode
		LastOp int
		Max    int
	}
	stack := []StackFrame{
		{
			Node:   root,
			LastOp: MoveLeft,
			Max:    0,
		},
	}
	for len(stack) > 0 {
		sf := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if sf.Node != nil {
			if sf.Max > max {
				max = sf.Max
			}
			switch sf.LastOp {
			case MoveLeft:
				if sf.Node.Left != nil {
					stack = append(stack, StackFrame{
						Node:   sf.Node.Left,
						LastOp: MoveRight,
						Max:    sf.Max + 1,
					})
				}
				if sf.Node.Right != nil {
					stack = append(stack, StackFrame{
						Node:   sf.Node.Right,
						LastOp: MoveLeft,
						Max:    1,
					})
				}
			case MoveRight:
				if sf.Node.Right != nil {
					stack = append(stack, StackFrame{
						Node:   sf.Node.Right,
						LastOp: MoveLeft,
						Max:    sf.Max + 1,
					})
				}
				if sf.Node.Left != nil {
					stack = append(stack, StackFrame{
						Node:   sf.Node.Left,
						LastOp: MoveRight,
						Max:    1,
					})
				}
			}
		}
	}
	return max
}
