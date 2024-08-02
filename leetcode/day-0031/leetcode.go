package leetcode

/*
	437: Path Sum III

	Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values
	along the path equals targetSum.

	The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from
	parent nodes to child nodes).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := countPathsWithSum(root, targetSum)
	count += pathSum(root.Left, targetSum)
	count += pathSum(root.Right, targetSum)
	return count
}

func countPathsWithSum(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val == targetSum {
		count++
	}
	count += countPathsWithSum(node.Left, targetSum-node.Val)
	count += countPathsWithSum(node.Right, targetSum-node.Val)
	return count
}

func insertLevelOrder(a []interface{}, root *TreeNode, i int, n int) *TreeNode {
	if i < n {
		var temp *TreeNode
		if a[i] != nil {
			temp = &TreeNode{
				Val: a[i].(int),
			}
		}
		root = temp
		if root != nil {
			root.Left = insertLevelOrder(a, root.Left, 2*i+1, n)
			root.Right = insertLevelOrder(a, root.Right, 2*i+2, n)
		}
	}
	return root
}

func pathSum_optimized(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	prefixSum := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSum)
}

func dfs(node *TreeNode, currSum, targetSum int, prefixSum map[int]int) int {
	if node == nil {
		return 0
	}

	currSum += node.Val
	count := prefixSum[currSum-targetSum]
	prefixSum[currSum]++

	count += dfs(node.Left, currSum, targetSum, prefixSum)
	count += dfs(node.Right, currSum, targetSum, prefixSum)

	prefixSum[currSum]--

	return count
}
