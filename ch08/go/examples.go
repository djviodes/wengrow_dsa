package ch08

func NoHashIsSubset(list1 []int, list2 []int) bool {
	largerList := []int{}
	smallerList := []int{}

	if len(list1) > len(list2) {
		largerList = list1
		smallerList = list2
	} else {
		largerList = list2
		smallerList = list1
	}

	for i := range len(smallerList) {
		foundMatch := false

		for j := range len(largerList) {
			if smallerList[i] == largerList[j] {
				foundMatch = true
				break
			}
		}

		if foundMatch == false {
			return false
		}
	}

	return true
}

func HashIsSubset(list1 []int, list2 []int) bool {
	largerList := []int{}
	smallerList := []int{}
	hashTable := map[int]bool{}

	if len(list1) > len(list2) {
		largerList = list1
		smallerList = list2
	} else {
		largerList = list2
		smallerList = list1
	}

	for _, num := range largerList {
		hashTable[num] = true
	}

	for _, num := range smallerList {
		if !hashTable[num] {
			return false
		}
	}

	return true
}
