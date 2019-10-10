package main

import (
	"fmt"
)

func sweep(numbers []int) {

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
		}
	}
}
func bubbleSort(numbers []int) {
	numbersLength := len(numbers)
	for i := 0; i < numbersLength; i++ {
		sweep(numbers)
	}
}

func main() {
	numbers := []int{2, 5, 6, 12, 0, 3, 7, 34, 21, 8}
	fmt.Println("Numbers: ", numbers)

	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)

}
