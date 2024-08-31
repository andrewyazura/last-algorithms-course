package quick_sort

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	pivot_i := 0
	pivot := array[pivot_i]

	for i := 1; i < len(array); i++ {
		if array[i] < pivot {
			array[pivot_i] = array[i]
			array[i] = array[pivot_i+1]
			pivot_i++
			array[pivot_i] = pivot
		}
	}

	QuickSort(array[:pivot_i])
	QuickSort(array[pivot_i+1:])
}

func QuickSort2(array []int) {
	qs(array, 0, len(array))
}

func qs(array []int, lo int, hi int) {
	if lo == hi {
		return
	}

	pivot_i := partition(array, lo, hi)
	qs(array, lo, pivot_i)
	qs(array, pivot_i+1, hi)
}

func partition(array []int, lo int, hi int) int {
	// array[i] where i <= pivot_i is less than pivot_i
	pivot_i := lo - 1
	pivot := array[hi-1]

	for i := lo; i < hi-1; i++ {
		if array[i] <= pivot {
			pivot_i++
			array[i], array[pivot_i] = array[pivot_i], array[i]
		}
	}

	pivot_i++

	array[hi-1] = array[pivot_i]
	array[pivot_i] = pivot
	return pivot_i
}
