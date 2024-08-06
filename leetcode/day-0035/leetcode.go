package leetcode

/*
	199. Binary Tree Right Side View
	Given the root of a binary ree, imagine yourself standing on the right side of it, return
	the values of nodes you can see ordered from top to bottom.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	q := []*TreeNode{
		root,
	}
	result := []int{}
	if root == nil {
		return result
	}
	var cur int
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]
			if n != nil {
				cur = n.Val
				if n.Left != nil {
					cur = n.Val
					q = append(q, n.Left)
				}
				if n.Right != nil {
					cur = n.Val
					q = append(q, n.Right)
				}
			}
		}
		result = append(result, cur)
	}
	return result
}

type Solution struct {
	Rightside []int
}

func (s *Solution) helper(node *TreeNode, level int) {
	if level == len(s.Rightside) {
		s.Rightside = append(s.Rightside, node.Val)
	}

	if node.Right != nil {
		s.helper(node.Right, level+1)
	}
	if node.Left != nil {
		s.helper(node.Left, level+1)
	}
}

func rightSideView_dfs(root *TreeNode) []int {
	s := Solution{
		Rightside: []int{},
	}
	if root == nil {
		return s.Rightside
	}
	s.helper(root, 0)
	return s.Rightside
}
