package ch01

import "fmt"

var groceryList = []string{"apples", "bananas", "cucumbers", "dates", "elderberries"}

func SearchForValue(item string) int {
	counter := 0

	for i, value := range groceryList {
		counter++
		if value == item {
			fmt.Println("Steps: ", counter)
			return i
		}
	}

	fmt.Println("Steps: ", counter)
	return -1
}

func InsertAtBeginning(item string) {
	groceryList = append(groceryList, "")
	counter := 0

	for i := len(groceryList) - 1; i > 0; i-- {
		counter++
		groceryList[i] = groceryList[i-1]
	}

	groceryList[0] = item
	counter++
	fmt.Println("Steps: ", counter)
}

func InsertAtEnd(item string) {
	groceryList = append(groceryList, item)

	fmt.Println("Steps: 1")
}

func DeleteAtBeginning() {
	counter := 0

	for i := 0; i < len(groceryList)-1; i++ {
		counter++
		groceryList[i] = groceryList[i+1]
	}

	groceryList = groceryList[:len(groceryList)-1]
	counter++
	fmt.Println("Steps: ", counter)
}

func DeleteAtEnd() {
	groceryList = groceryList[:len(groceryList)-1]

	fmt.Println("Steps: 1")
}
