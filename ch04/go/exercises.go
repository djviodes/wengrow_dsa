package ch04

import "fmt"

func QuadraticGreatestNumber(list [6]int) int {
	counter := 0

	for _, i := range list {
		isIValTheGreatest := true

		for _, j := range list {
			counter++

			if j > i {
				isIValTheGreatest = false
			}
		}

		if isIValTheGreatest {
			fmt.Println("Greatest Number: ", i)
			fmt.Println("Steps: ", counter)
			return counter
		}
	}

	return counter
}

func LinearGreatestNumber(list [6]int) int {
	greatestNumber := list[0]
	counter := 0

	for i := range list {
		counter++

		if list[i] > greatestNumber {
			greatestNumber = list[i]
		}
	}

	fmt.Println("Greatest Number: ", greatestNumber)
	fmt.Println("Steps: ", counter)
	return counter
}
