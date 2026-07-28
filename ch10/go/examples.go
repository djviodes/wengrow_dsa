package ch10

import "fmt"

func Countdown(number int) int {
	counter := 0

	for i := number; i > 0; i-- {
		counter++
		fmt.Println(i)
	}

	fmt.Println("Steps: ", counter)
	return counter
}

func RecursiveCountdown(number int, counter int) int {
	if number > 0 {
		fmt.Println(number)
		counter++
		number--
		return RecursiveCountdown(number, counter)
	} else {
		fmt.Println("Steps: ", counter)
		return counter
	}
}

func Factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * Factorial(number-1)
	}
}
