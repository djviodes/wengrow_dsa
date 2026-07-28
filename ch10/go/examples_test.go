package ch10

import "testing"

func TestFactorial(t *testing.T) {
	if result := Factorial(5); result != 120 {
		t.Errorf("expected 120, got %d", result)
	}

	if result := Factorial(1); result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

func ExampleCountdown() {
	Countdown(3)
	// Output: 3
	// 2
	// 1
	// Steps:  3
}

func ExampleRecursiveCountdown() {
	RecursiveCountdown(3, 0)
	// Output: 3
	// 2
	// 1
	// Steps:  3
}
