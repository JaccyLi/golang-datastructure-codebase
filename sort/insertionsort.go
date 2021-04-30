package sort

// InsertionSort
// Insertion sort works similarly as we sort cards in our hand in a card game.
/*
We assume that the first card is already sorted then,
we select an unsorted card. If the unsorted card is
greater than the card in hand, it is placed on the
right otherwise, to the left. In the same way, other
unsorted cards are taken and put at their right place.
*/

//func InsertionSort(unsort []int) []int {
//	sorted := []int{unsort[0]}
//	for i := 1; i < len(unsort); i++ {
//		for j := 1; j < len(sorted); j++ {
//			if unsort[i] > sorted[j] {
//				sorted[j+1] = unsort[i]
//			} else {
//				sorted[j-1] = unsort[i]
//			}
//		}
//	}
//	return sorted
//}
