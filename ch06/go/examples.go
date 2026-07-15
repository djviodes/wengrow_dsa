package ch06

import "fmt"

func InsertionSort(list [5]int) int {
	counter := 0

	for i := 1; i < len(list); i++ {
		counter++
		temp := list[i]
		position := i - 1

		for position >= 0 {
			counter++

			if list[position] > temp {
				counter++
				list[position+1] = list[position]
				position--
			} else {
				break
			}
		}

		counter++
		list[position+1] = temp
	}

	fmt.Println("Sorted List: ", list)
	fmt.Println("Steps: ", counter)
	return counter
}

func NonbreakIntersection(firstList [4]int, secondList [4]int) int {
	result := []int{}
	counter := 0

	for i := range firstList {
		for j := range secondList {
			counter++

			if firstList[i] == secondList[j] {
				counter++
				result = append(result, firstList[i])
			}
		}
	}

	fmt.Println("Result: ", result)
	fmt.Println("Steps: ", counter)
	return counter
}

func BreakIntersection(firstList [4]int, secondList [4]int) int {
	result := []int{}
	counter := 0

	for i := range firstList {
		for j := range secondList {
			counter++

			if firstList[i] == secondList[j] {
				counter++
				result = append(result, firstList[i])
				break
			}
		}
	}

	fmt.Println("Result: ", result)
	fmt.Println("Steps: ", counter)
	return counter
}
