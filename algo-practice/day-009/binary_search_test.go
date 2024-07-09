package practice

import "testing"

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	idx, count := BinarySearch(list, 8)
	if idx != 7 {
		t.FailNow()
	}
	if count != 1 {
		t.FailNow()
	}
	idx, count = BinarySearch(list, 4)
	if idx != 3 {
		t.FailNow()
	}
	if count != 2 {
		t.FailNow()
	}
}
