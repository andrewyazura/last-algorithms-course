package bubble_sort

import (
  "testing"
  "slices"
  "math/rand"
)

func TestBubbleSort(t *testing.T) {
  array := rand.Perm(10_000)

  BubbleSort(array)

  if !slices.IsSorted(array) {
      t.Errorf("array is not sorted")
  }
}
