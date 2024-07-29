package leetcode

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// find max
	list := list.New()
	max := 0
	node := head
	for node != nil {
		list.PushBack(node.Val)
		node = node.Next
	}
	// remove front and back
	for list.Len() > 0 {
		v1, _ := list.Remove(list.Front()).(int)
		v2, _ := list.Remove(list.Back()).(int)
		sum := v1 + v2
		if sum > max {
			max = sum
		}
	}
	return max
}

func pairSum_reverse(head *ListNode) int {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	res := 0
	for slow != nil {
		sum := prev.Val + slow.Val
		if sum > res {
			res = sum
		}
		prev = prev.Next
		slow = slow.Next
	}
	return res
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
