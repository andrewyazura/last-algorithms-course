package bubble_sort

func BubbleSort(array []int) {
  for i := range array {
    for j := 0; j < len(array) - i - 1; j++ {
      if array[j] > array[j+1] {
        array[j], array[j+1] = array[j+1], array[j]
      }
    }
  }
}
