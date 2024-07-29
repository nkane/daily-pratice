package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDeleteMiddle(t *testing.T) {
	list := &ListNode{
		Val: 1,
	}
	list.Next = &ListNode{
		Val: 3,
	}
	list.Next.Next = &ListNode{
		Val: 4,
	}
	list.Next.Next.Next = &ListNode{
		Val: 7,
	}
	list.Next.Next.Next.Next = &ListNode{
		Val: 1,
	}
	list.Next.Next.Next.Next.Next = &ListNode{
		Val: 2,
	}
	list.Next.Next.Next.Next.Next.Next = &ListNode{
		Val: 6,
	}
	expected := getList(list)
	assert.DeepEqual(t, expected, []int{1, 3, 4, 7, 1, 2, 6})

	list = deleteMiddle(list)
	expected = []int{1, 3, 4, 1, 2, 6}
	got := getList(list)
	assert.DeepEqual(t, expected, got)

	list = &ListNode{
		Val: 1,
	}
	expected = getList(list)
	assert.DeepEqual(t, expected, []int{1})
	list = deleteMiddle(list)
	expected = []int{}
	got = getList(list)
	assert.DeepEqual(t, expected, got)
}
