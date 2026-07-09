package ch01

func resetGroceryList() {
	groceryList = []string{"apples", "bananas", "cucumbers", "dates", "elderberries"}
}

func ExampleSearchForValue() {
	resetGroceryList()
	searchForValue("elderberries")
	// Output: Steps:  5
}

func ExampleInsertAtBeginning() {
	resetGroceryList()
	insertAtBeginning("figs")
	// Output: Steps:  6
}

func ExampleInsertAtEnd() {
	resetGroceryList()
	insertAtEnd("pomegranates")
	// Output: Steps: 1
}

func ExampleDeleteAtBeginning() {
	resetGroceryList()
	deleteAtBeginning()
	// Output: Steps:  5
}

func ExampleDeleteAtEnd() {
	resetGroceryList()
	deleteAtEnd()
	// Output: Steps: 1
}
