package main

import "testing"

func TestBubbleEmpty(t *testing.T) {
	res := sortBubble(nil)
	if res != nil {
		t.Error("empty slice is not nil")
	}
}

func TestBubbleSuccess(t *testing.T) {
	res := sortBubble([]int{2, 3, 1})
	if !sliceEqualHelper(res, []int{1, 2, 3}) {
		t.Error("empty slice is not nil")
	}
}

func TestBubbleNoSortNeed(t *testing.T) {
	res := sortBubble([]int{1, 3, 5})
	if !sliceEqualHelper(res, []int{1, 3, 5}) {
		t.Error("empty slice is not nil")
	}
}

func sliceEqualHelper(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for id := range a {
		if a[id] != b[id] {
			return false
		}
	}
	return true
}
