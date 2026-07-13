package ch04

func ExampleBubbleSort() {
	BubbleSort([5]int{4, 2, 7, 1, 3})
	// Output: Steps:  16
	// Sorted List:  [1 2 3 4 7]
}

func ExampleQuadraticHasDuplicateValues() {
	QuadraticHasDuplicateValues([6]int{1, 5, 3, 9, 1, 4})
	// Output: Has Duplicate Values
	// Steps:  6
}

func ExampleLinearHasDuplicateValues() {
	LinearHasDuplicateValues([6]int{1, 5, 3, 9, 1, 4})
	// Output: Has Duplicate Values
	// Steps:  5
}
