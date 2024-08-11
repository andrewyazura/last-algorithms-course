package binary_search

func BinarySearch(haystack []int, needle int) int {
  lo := 0
  hi := len(haystack)

  for lo < hi {
    mi := lo + (hi - lo) / 2

    switch {
    case haystack[mi] == needle:
        return mi
    case haystack[mi] < needle:
        lo = mi + 1
    case haystack[mi] > needle:
        hi = mi
    }
  }

  return -1
}
