package ch07

import (
	"fmt"
	"strconv"
)

func AverageOfEvenNumbers(list []int) int {
	sum := 0.0
	countOfEvenNumbers := 0
	counter := 0

	for _, number := range list {
		counter++

		if number%2 == 0 {
			counter++
			sum += float64(number)
			countOfEvenNumbers++
		}
	}

	fmt.Println("Average: ", sum/float64(countOfEvenNumbers))
	fmt.Println("Steps: ", counter)
	return counter
}

func WordBuilder(list []string) int {
	collection := []string{}
	counter := 0

	for i := range len(list) {
		for j := range len(list) {
			counter++

			if i != j {
				counter++
				collection = append(collection, list[i]+list[j])
			}
		}
	}

	fmt.Println("Words: ", collection)
	fmt.Println("Steps: ", counter)
	return counter
}

func ListSample(list []int) int {
	first := list[0]
	middle := list[int(len(list)/2)]
	last := list[len(list)-1]
	counter := 0

	counter++

	fmt.Println("List Sample: ", []int{first, middle, last})
	fmt.Println("Steps: ", counter)
	return counter
}

func AverageCelcius(fahrenheitReadings []float32) int {
	celciusNumbers := []float32{}
	counter := 0

	for _, fahrenheitReading := range fahrenheitReadings {
		counter++

		celciusConversion := (fahrenheitReading - 32) / 1.8
		counter++

		celciusNumbers = append(celciusNumbers, celciusConversion)
		counter++
	}

	sum := 0.0

	for _, celciusNumber := range celciusNumbers {
		counter++

		sum += float64(celciusNumber)
		counter++
	}

	fmt.Println("Mean Average: ", (sum / float64(len(celciusNumbers))))
	fmt.Println("Steps: ", counter)
	return counter
}

func MarkInventory(clothingItems []string) int {
	clothingOptions := []string{}
	counter := 0

	for _, item := range clothingItems {
		for size := 1; size <= 5; size++ {
			counter++
			clothingOptions = append(clothingOptions, item+" Size: "+strconv.Itoa(size))
		}
	}

	fmt.Println("Clothing Options: ", clothingOptions)
	fmt.Println("Steps: ", counter)
	return counter
}

func CountOnes(matrix [][]int) int {
	onesCount := 0
	counter := 0

	for _, vector := range matrix {
		for _, num := range vector {
			counter++

			if num == 1 {
				onesCount++
			}
		}
	}

	fmt.Println("Ones Count: ", onesCount)
	fmt.Println("Steps: ", counter)
	return counter
}

func IsPalindrome(str string) int {
	leftIndex := 0
	rightIndex := len(str) - 1
	counter := 0

	for float64(leftIndex) < float64(len(str))/2 {
		counter++

		if str[leftIndex] != str[rightIndex] {
			fmt.Println("Is Palindrome: ", false)
			fmt.Println("Steps: ", counter)
			return counter
		}

		leftIndex++
		rightIndex--
	}

	fmt.Println("Is Palindrome: ", true)
	fmt.Println("Steps: ", counter)
	return counter
}

func TwoNumberProducts(list []int) int {
	products := []int{}
	counter := 0

	for i := range len(list) {
		for j := range len(list) {
			counter++
			products = append(products, list[i]*list[j])
		}
	}

	fmt.Println("Products: ", products)
	fmt.Println("Steps: ", counter)
	return counter
}
