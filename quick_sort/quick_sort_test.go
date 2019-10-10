package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 3, 4, 5, 17, 1}, []int{1, 2, 3, 4, 5, 17}},
		{[]int{13, 2, 4, 21}, []int{2, 4, 13, 21}},
		{[]int{2, 5, 1, 12, 11}, []int{1, 2, 5, 11, 12}},
	}
	for _, test := range tests {
		QuickSort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("slice %v should be sorted as %v", test.expected, test.input)
		}
	}
}
