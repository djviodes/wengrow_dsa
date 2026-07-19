package ch07

func ExampleAverageOfEvenNumbers() {
	AverageOfEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// Output: Average:  6
	// Steps:  15
}

func ExampleWordBuilder() {
	WordBuilder([]string{"a", "b", "c", "d"})
	// Output: Words:  [ab ac ad ba bc bd ca cb cd da db dc]
	// Steps:  28
}

func ExampleListSample() {
	ListSample([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// Output: List Sample:  [0 5 10]
	// Steps:  1
}

func ExampleAverageCelcius() {
	AverageCelcius([]float32{82.0, 81.2, 85.3, 84.0, 78.6, 79.9, 80.7, 75.1})
	// Output: Mean Average:  27.13888955116272
	// Steps:  40
}

func ExampleMarkInventory() {
	MarkInventory([]string{"Purple Shirt", "Green Shirt"})
	// Output: Clothing Options:  [Purple Shirt Size: 1 Purple Shirt Size: 2 Purple Shirt Size: 3 Purple Shirt Size: 4 Purple Shirt Size: 5 Green Shirt Size: 1 Green Shirt Size: 2 Green Shirt Size: 3 Green Shirt Size: 4 Green Shirt Size: 5]
	// Steps:  10
}

func ExampleCountOnes() {
	CountOnes([][]int{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 0, 1, 1},
		{1, 1, 0, 0},
	})
	// Output: Ones Count:  8
	// Steps:  16
}

func ExampleIsPalindrome() {
	IsPalindrome("racecar")
	// Output: Is Palindrome:  true
	// Steps:  4
}

func ExampleTwoNumberProducts() {
	TwoNumberProducts([]int{1, 2, 3, 4, 5})
	// Output: Products:  [1 2 3 4 5 2 4 6 8 10 3 6 9 12 15 4 8 12 16 20 5 10 15 20 25]
	// Steps:  25
}
