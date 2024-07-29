package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	stack := []*ListNode{}
	node := head
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	if len(stack) == 0 {
		return nil
	}
	// build list
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}
	stack[0].Next = nil
	return stack[len(stack)-1]
}

func getList(head *ListNode) []int {
	l := []int{}
	node := head
	for node != nil {
		l = append(l, node.Val)
		node = node.Next
	}
	return l
}
