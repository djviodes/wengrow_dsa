package ch03

import "fmt"

var things = []string{"apples", "baboons", "cribs", "dulcimers"}

func PrintThings() int {
	counter := 0

	for _, thing := range things {
		counter++
		fmt.Println("Here's a thing: ", thing)
	}

	fmt.Println("Steps: ", counter)
	return counter
}

func IsPrime(number int) (bool, int) {
	counter := 0

	for i := 2; i < number; i++ {
		counter++

		if number%i == 0 {
			fmt.Println("Steps: ", counter)
			return false, counter
		}
	}

	fmt.Println("Steps: ", counter)
	return true, counter
}
