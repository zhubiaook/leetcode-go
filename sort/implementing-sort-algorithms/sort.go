package main

import "cmp"

// InsertionSort is a simple sorting algorithm that works similarly to the way you sort playing cards in your hands.
// The array is virtually split into a sorted and an unsorted part.
// Values from the unsorted part are picked and placed in the correct position in the sorted part.
func InsertionSort[T cmp.Ordered](items []T) {
	if len(items) <= 1 {
		return
	}

	// Insert the unsorted element item[i] into the sorted elements until all elements are sorted
	for i := 1; i < len(items); i++ {
		for j := i - 1; j >= 0; j-- {
			if items[j+1] >= items[j] {
				break
			}
			items[j+1], items[j] = items[j], items[j+1]
		}
	}
}

// BubbleSort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements
// if they are in the wrong order.
// This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.
func BubbleSort[T cmp.Ordered](itmes []T) {
	for i := 0; i < len(itmes)-1; i++ {
		for j := 0; j < len(itmes)-i-1; j++ {
			if itmes[j] > itmes[j+1] {
				itmes[j], itmes[j+1] = itmes[j+1], itmes[j]
			}
		}
	}
}

// SelectionSort is a sorting algorithm that selects the smallest element from the array
func SelectionSort[T cmp.Ordered](items []T) {
	// Find the smallest element in the list
	for i := 0; i < len(items)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			items[i], items[minIndex] = items[minIndex], items[i]
		}
	}
}
