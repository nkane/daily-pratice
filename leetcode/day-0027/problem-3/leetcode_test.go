package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestOddEvenList(t *testing.T) {
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

	list = oddEvenList(list)
	expected = []int{1, 3, 5, 2, 4}
	got := getList(list)
	assert.DeepEqual(t, expected, got)
}
