package binary_search

import (
  "errors"
)

func BinarySearch(haystack []int, needle int) (int, error) {
  lo := 0
  hi := len(haystack)

  for lo < hi {
    mi := lo + (hi - lo) / 2

    switch {
    case haystack[mi] == needle:
        return mi, nil
    case haystack[mi] < needle:
        lo = mi + 1
    case haystack[mi] > needle:
        hi = mi
    }
  }

  return 0, errors.New("needle not in haystack")
}
