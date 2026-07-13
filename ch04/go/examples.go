package ch04

import "fmt"

var unsorted_list = []int{4, 2, 7, 1, 3}

func BubbleSort(list [5]int) int {
	unsorted_list_index := len(unsorted_list) - 1
	is_sorted := false
	counter := 0

	for !is_sorted {
		is_sorted = true

		for i := range unsorted_list_index {
			counter++

			if list[i] > list[i+1] {
				counter++
				list[i], list[i+1] = list[i+1], list[i]
				is_sorted = false
			}
		}

		unsorted_list_index--
	}

	fmt.Println("Steps: ", counter)
	fmt.Println("Sorted List: ", list)
	return counter
}

func QuadraticHasDuplicateValues(list [6]int) int {
	counter := 0

	for i, _ := range list {
		counter++

		for j, _ := range list {
			counter++

			if i != j && list[i] == list[j] {
				fmt.Println("Has Duplicate Values")
				fmt.Println("Steps: ", counter)
				return counter
			}
		}
	}

	fmt.Println("Does Not Have Duplicate Values")
	fmt.Println("Steps: ", counter)
	return counter
}

func LinearHasDuplicateValues(list [6]int) int {
	existingValues := [10]int{}
	counter := 0

	for i, _ := range list {
		counter++

		if existingValues[list[i]] == 1 {
			fmt.Println("Has Duplicate Values")
			fmt.Println("Steps: ", counter)
			return counter
		} else {
			existingValues[list[i]] = 1
		}
	}

	fmt.Println("Does Not Have Duplicate Values")
	fmt.Println("Steps: ", counter)
	return counter
}
