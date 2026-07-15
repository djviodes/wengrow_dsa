package ch06

func ExampleInsertionSort() {
	InsertionSort([5]int{4, 2, 7, 1, 3})
	// Output: Sorted List:  [1 2 3 4 7]
	// Steps:  22
}

func ExampleNonbreakIntersection() {
	NonbreakIntersection([4]int{3, 1, 4, 2}, [4]int{4, 5, 3, 6})
	// Output: Result:  [3 4]
	// Steps:  18
}

func ExampleBreakIntersection() {
	BreakIntersection([4]int{3, 1, 4, 2}, [4]int{4, 5, 3, 6})
	// Output: Result:  [3 4]
	// Steps:  14
}
