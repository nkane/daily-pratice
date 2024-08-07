package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 1
	}
	q := []*TreeNode{
		root,
	}
	maxSum := root.Val
	maxDepth := 1
	depth := 1
	for len(q) > 0 {
		l := len(q)
		tmpSum := 0
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]
			tmpSum += n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if tmpSum > maxSum {
			maxSum = tmpSum
			maxDepth = depth
		}
		depth++
	}
	return maxDepth
}
