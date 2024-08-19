package quick_sort

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	pivot_i := 0
	pivot := array[pivot_i]
	i := 1

	for i < len(array) {
		if array[i] < pivot {
			array[pivot_i] = array[i]
			array[i] = array[pivot_i+1]
			pivot_i++
			array[pivot_i] = pivot
		}

		i++
	}

	QuickSort(array[:pivot_i])
	QuickSort(array[pivot_i+1:])
}
