package quick_sort

import (
  "testing"
  "slices"
  "math/rand"
)

func TestQuickSort(t *testing.T) {
  array := rand.Perm(1_000_000)
  
  QuickSort(array)

  if !slices.IsSorted(array) {
    t.Errorf("array is not sorted")
  }
}
