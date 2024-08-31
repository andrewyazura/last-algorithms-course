package quick_sort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := rand.Perm(1_000_000)

	QuickSort(array)

	if !slices.IsSorted(array) {
		t.Errorf("array is not sorted")
	}
}

func TestQuickSort2(t *testing.T) {
	array := rand.Perm(1_000_000)

	QuickSort2(array)

	if !slices.IsSorted(array) {
		t.Errorf("array is not sorted")
	}
}
