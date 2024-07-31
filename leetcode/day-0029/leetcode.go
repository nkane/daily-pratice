package leetcode

/*
	872. Leaf-Similar Trees
	Consider all the leaves of a binary tree, from left to right order, the values of those leaves from
	a leaf value sequence.

	Two binary trees are considered leaf-similar if their leaf values sequence is the same. Return true
	if and only if the two given trees with head hodes root1 and root2 are leaf-similar
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := DFSReturnLeaves(root1)
	leaves2 := DFSReturnLeaves(root2)
	if len(leaves1) == len(leaves2) {
		for i := 0; i < len(leaves1); i++ {
			if leaves1[i] != leaves2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func DFSReturnLeaves(root *TreeNode) []int {
	leaves := []int{}
	if root == nil {
		return leaves
	}
	stack := []*TreeNode{
		root,
	}
	for len(stack) > 0 {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if element != nil {
			if element.Right == nil && element.Left == nil {
				leaves = append(leaves, element.Val)
				continue
			}
			stack = append(stack, element.Left)
			stack = append(stack, element.Right)
		}
	}
	return leaves
}

func BFSReturnLeaves(root *TreeNode) []int {
	leaves := []int{}
	if root == nil {
		return leaves
	}
	q := []*TreeNode{
		root,
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.Right == nil && n.Left == nil {
			leaves = append(leaves, n.Val)
			continue
		}
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return leaves
}
