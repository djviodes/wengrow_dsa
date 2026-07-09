package ch02

func resetOrderedArray() {
	orderedArray = []int{2, 4, 6, 7, 9, 12, 14, 17, 21}
}

func ExampleLinearSearch() {
	resetOrderedArray()
	LinearSearch(14)
	// Output: Steps:  7
}

func ExampleBinarySearch() {
	resetOrderedArray()
	BinarySearch(14)
	// Output: Steps:  2
}

func ExampleNaiveInsertion() {
	resetOrderedArray()
	NaiveInsertion(12)
	// Output: Steps:  11
}

func ExampleBinaryInsertion() {
	resetOrderedArray()
	BinaryInsertion(12)
	// Output: Steps:  7
}
