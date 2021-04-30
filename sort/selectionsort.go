package sort

// SelectSort
// assend: find minimum number, and put it at beginning of unsorted part
func SelectSortAssend(unsort []int) []int {
	for i := 0; i < len(unsort); i++ {
		minIndex := i
		for j := i + 1; j < len(unsort); j++ {
			if unsort[j] < unsort[minIndex] {
				// find unsorted part's minmum value's index
				minIndex = j
			}
		}
		// swap minmum index with first one(unsorted part)
		unsort[minIndex], unsort[i] = unsort[i], unsort[minIndex]
	}
	return unsort
}

// descend:
func SelectSortDescend(unsort []int) []int {
	for i := 0; i < len(unsort); i++ {
		minIndex := i
		for j := i + 1; j < len(unsort); j++ {
			if unsort[j] > unsort[minIndex] {
				// find unsorted part's minmum value's index
				minIndex = j
			}
		}
		// swap minmum index with first one(unsorted part)
		unsort[minIndex], unsort[i] = unsort[i], unsort[minIndex]
	}
	return unsort
}
