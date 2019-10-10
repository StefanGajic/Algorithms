package main

import (
	"fmt"
)

func bubbleSort(numbers []int) {
	for {
		swapped := false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				swapped = true
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	numbers := []int{5, 2, 4, 6, 7, 8, 9, 10, 11, 13, 19}
	fmt.Println("Numbers: ", numbers)

	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)

}
