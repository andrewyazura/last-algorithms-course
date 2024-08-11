package bubble_sort

import (
  "testing"
)

func TestBubbleSort(t *testing.T) {
  arr := make([]int, 10)
  for i := range arr {
    arr[9 - i] = i
  }

  wantArr := make([]int, 10)
  for i := range wantArr {
    wantArr[i] = i
  }

  BubbleSort(arr)

  for i := range arr {
    if arr[i] != wantArr[i] {
      t.Errorf("BubbleSort(%v) != %v", arr, wantArr)
    }
  }
}
