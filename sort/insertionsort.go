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

// 3,4,2,1,4,6,8,3
// 3   4,2,1,4,6,8,3
// 3,4  2,1,4,6,8,3
// 2,3,4 1,4,

//func InsertionSort(unsort []int) []int {
//	for i := 1; i < len(unsort); i++ {
//		tmp := []int{unsort[i]}
//		for j := i - 1; j > 0; j-- {
//			if tmp[j] > unsort[i] {
//
//			}
//
//		}
//
//	}
//
//	return []int{}
//}
//
