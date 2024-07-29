package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestReverseList(t *testing.T) {
	list := &ListNode{
		Val: 1,
	}
	list.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next = &ListNode{
		Val: 3,
	}
	list.Next.Next.Next = &ListNode{
		Val: 4,
	}
	list.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	expected := getList(list)
	assert.DeepEqual(t, expected, []int{1, 2, 3, 4, 5})

	list = reverseList(list)
	expected = []int{5, 4, 3, 2, 1}
	got := getList(list)
	assert.DeepEqual(t, expected, got)
}
