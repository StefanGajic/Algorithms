package main

import (
	"fmt"
)

func partition(a []int, low, high int) int {
	p := a[high]
	for j := low; j < high; j++ {
		if a[j] < p {
			a[j], a[low] = a[low], a[j]
			low++
		}
	}
	a[low], a[high] = a[high], a[low]
	return low
}

func quickSort(a []int, low, high int) {
	if low > high {
		return
	}
	pivot := partition(a, low, high)
	quickSort(a, low, pivot-1)
	quickSort(a, pivot+1, high)
}

func QuickSort(q []int) {
	quickSort(q, 0, len(q)-1)
}

func main() {
	numbers := []int{2, 5, 6, 12, 0, 3, 7, 34, 21, 8}
	fmt.Println("Numbers: ", numbers)

	QuickSort(numbers)
	fmt.Println("After Sorting: ", numbers)
}
