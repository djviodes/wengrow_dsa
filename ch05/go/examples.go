package ch05

import "fmt"

func SelectionSort(list [5]int) int {
	counter := 0

	for i := range len(list) {
		lowestNumberIndex := i

		for j := i + 1; j < len(list); j++ {
			counter++

			if list[j] < list[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		if lowestNumberIndex != i {
			counter++
			temp := list[i]
			list[i] = list[lowestNumberIndex]
			list[lowestNumberIndex] = temp
		}
	}

	fmt.Println("Sorted List: ", list)
	fmt.Println("Steps: ", counter)
	return counter
}

func PrintNumbersVersionOne(upperLimit int) int {
	number := 2
	counter := 0

	for number <= upperLimit {
		counter++

		// If number is even, print it:
		if number%2 == 0 {
			fmt.Println(number)
		}

		number++
	}

	return counter
}

func PrintNumbersVersionTwo(upperLimit int) int {
	number := 2
	counter := 0

	for number <= upperLimit {
		counter++
		fmt.Println(number)

		// Increase number by 2, which, by definition
		// is the next even number:
		number += 2
	}

	return counter
}
