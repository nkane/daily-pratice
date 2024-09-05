package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var fast, slow *ListNode
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
