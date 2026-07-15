package ch06

func ExampleWorstCaseContainsX() {
	WorstCaseContainsX("TestX")
	// Output: Found X:  true
	// Steps:  5
}

func ExampleBestAndAverageCaseContainsX_best() {
	BestAndAverageCaseContainsX("XTest")
	// Output: Found X:  true
	// Steps:  1
}

func ExampleBestAndAverageCaseContainsX_average() {
	BestAndAverageCaseContainsX("TeXst")
	// Output: Found X:  true
	// Steps:  3
}
