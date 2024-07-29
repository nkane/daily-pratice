package leetcode

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPairSums(t *testing.T) {
	list := &ListNode{
		Val: 4,
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
	got := pairSum(list)
	expected := 8
	assert.DeepEqual(t, expected, got)
}
