package ch01

func resetGroceryList() {
	groceryList = []string{"apples", "bananas", "cucumbers", "dates", "elderberries"}
}

func ExampleSearchForValue() {
	resetGroceryList()
	SearchForValue("elderberries")
	// Output: Steps:  5
}

func ExampleInsertAtBeginning() {
	resetGroceryList()
	InsertAtBeginning("figs")
	// Output: Steps:  6
}

func ExampleInsertAtEnd() {
	resetGroceryList()
	InsertAtEnd("pomegranates")
	// Output: Steps: 1
}

func ExampleDeleteAtBeginning() {
	resetGroceryList()
	DeleteAtBeginning()
	// Output: Steps:  5
}

func ExampleDeleteAtEnd() {
	resetGroceryList()
	DeleteAtEnd()
	// Output: Steps: 1
}
