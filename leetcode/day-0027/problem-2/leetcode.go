package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// find length first
	l := 0
	node := head
	for node != nil {
		l++
		node = node.Next
	}
	if l == 1 {
		return nil
	}
	idx := int(l / 2)
	fmt.Printf("length: %d\n", l)
	fmt.Printf("idx: %d\n", idx)
	node = head
	prev := node
	for i := 0; i < idx; i++ {
		prev = node
		node = node.Next
	}
	prev.Next = node.Next
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
