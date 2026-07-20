package ch08

import "testing"

func TestNoHashIsSubset(t *testing.T) {
	if !NoHashIsSubset([]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}) {
		t.Errorf("expected [2, 4, 6] to be a subset of [1, 2, 3, 4, 5, 6]")
	}

	if NoHashIsSubset([]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 7}) {
		t.Errorf("expected [2, 4, 7] not to be a subset of [1, 2, 3, 4, 5, 6]")
	}
}

func TestHashIsSubset(t *testing.T) {
	if !HashIsSubset([]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}) {
		t.Errorf("expected [2, 4, 6] to be a subset of [1, 2, 3, 4, 5, 6]")
	}

	if HashIsSubset([]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 7}) {
		t.Errorf("expected [2, 4, 7] not to be a subset of [1, 2, 3, 4, 5, 6]")
	}
}
