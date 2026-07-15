package ch06

import "fmt"

func WorstCaseContainsX(str string) int {
	foundX := false
	counter := 0

	for i := range len(str) {
		counter++

		if string(str[i]) == "X" {
			foundX = true
		}
	}

	fmt.Println("Found X: ", foundX)
	fmt.Println("Steps: ", counter)
	return counter
}

func BestAndAverageCaseContainsX(str string) int {
	counter := 0

	for i := range len(str) {
		counter++

		if string(str[i]) == "X" {
			fmt.Println("Found X:  true")
			fmt.Println("Steps: ", counter)
			return counter
		}
	}

	fmt.Println("Found X:  false")
	fmt.Println("Steps: ", counter)
	return counter
}
