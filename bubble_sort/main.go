package main

import (
	"fmt"
)

func sweep(numbers []int) {
	numbersLength := len(numbers)
	firstIndex := 0
	secondIndex := 1

	for secondIndex < numbersLength {
		firstNumber := numbers[firstIndex]
		secondNumber := numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
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
