package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
  arr := make([]int, 10)
  for i := range arr {
    arr[i] = i
  }

  cases := []struct {
    in, want int
  } {
    {0, 0},
    {1, 1},
    {2, 2},
    {12, -1},
  }

  for _, c := range cases {
    got := BinarySearch(arr, c.in)

    if got != c.want {
      t.Errorf("BinarySearch(%d) == %d, expected %d", c.in, got, c.want)
    }
  }
}
