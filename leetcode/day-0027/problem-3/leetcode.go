package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddList := head
	evenList := head.Next
	evenHead := evenList
	for evenList != nil && evenList.Next != nil {
		oddList.Next = evenList.Next
		oddList = oddList.Next
		evenList.Next = oddList.Next
		evenList = evenList.Next
	}
	oddList.Next = evenHead
	return head
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
