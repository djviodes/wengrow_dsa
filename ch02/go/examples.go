package ch02

import "fmt"

var orderedArray = []int{2, 4, 6, 7, 9, 12, 14, 17, 21}

func LinearSearch(value int) (int, int) {
	counter := 0

	for i, v := range orderedArray {
		counter++
		if v == value {
			fmt.Println("Steps: ", counter)
			return i, counter
		}
	}

	fmt.Println("Steps: ", counter)
	return -1, counter
}

func BinarySearch(value int) (int, int) {
	lowerBound := 0
	upperBound := len(orderedArray) - 1
	counter := 0

	for lowerBound <= upperBound {
		counter++
		midpoint := (lowerBound + upperBound) / 2

		valueAtMidpoint := orderedArray[midpoint]

		if value == valueAtMidpoint {
			fmt.Println("Steps: ", counter)
			return midpoint, counter
		} else if value < valueAtMidpoint {
			upperBound = midpoint - 1
		} else {
			lowerBound = midpoint + 1
		}
	}

	fmt.Println("Steps: ", counter)
	return -1, counter
}

func NaiveInsertion(value int) int {
	counter := 0

	for i := 0; i < len(orderedArray); i++ {
		counter++
		if orderedArray[i] > value {
			orderedArray = append(orderedArray, 0)

			for j := len(orderedArray) - 1; j > i; j-- {
				counter++
				orderedArray[j] = orderedArray[j-1]
			}

			orderedArray[i] = value
			counter++

			fmt.Println("Steps: ", counter)
			return counter
		}
	}

	orderedArray = append(orderedArray, value)
	counter++

	fmt.Println("Steps: ", counter)
	return counter
}

func BinaryInsertion(value int) int {
	lowerBound := 0
	upperBound := len(orderedArray) - 1
	counter := 0

	for lowerBound <= upperBound {
		counter++
		midpoint := (lowerBound + upperBound) / 2

		valueAtMidpoint := orderedArray[midpoint]

		if value < valueAtMidpoint {
			upperBound = midpoint - 1
		} else {
			lowerBound = midpoint + 1
		}
	}

	orderedArray = append(orderedArray, 0)

	for i := len(orderedArray) - 1; i > lowerBound; i-- {
		counter++
		orderedArray[i] = orderedArray[i-1]
	}

	orderedArray[lowerBound] = value
	counter++

	fmt.Println("Steps: ", counter)
	return counter
}
