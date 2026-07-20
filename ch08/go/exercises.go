package ch08

func ListIntersection(list1 []int, list2 []int) []int {
	largerList := []int{}
	smallerList := []int{}
	intersections := []int{}
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
		if hashTable[num] {
			intersections = append(intersections, num)
		}
	}

	return intersections
}

func DuplicateChars(list []string) string {
	hashTable := map[string]bool{}

	for _, char := range list {
		if hashTable[char] {
			return char
		} else {
			hashTable[char] = true
		}
	}

	return "N/A"
}

func MissingLetters(iString string) rune {
	hashTable := map[rune]bool{'a': false, 'b': false, 'c': false, 'd': false, 'e': false,
		'f': false, 'g': false, 'h': false, 'i': false, 'j': false,
		'k': false, 'l': false, 'm': false, 'n': false, 'o': false,
		'p': false, 'q': false, 'r': false, 's': false, 't': false,
		'u': false, 'v': false, 'w': false, 'x': false, 'y': false,
		'z': false}

	for _, char := range iString {
		hashTable[char] = true
	}

	for key, value := range hashTable {
		if value == false {
			return key
		}
	}

	return '?'
}

func NonDuplicatedString(iString string) rune {
	runes := []rune(iString)
	lowestIndexSingleChar := 0
	hashTable := map[rune]int{}

	for _, char := range runes {
		if _, ok := hashTable[char]; !ok {
			hashTable[char] = 1
		} else {
			hashTable[char]++
		}
	}

	for lowestIndexSingleChar < len(runes) {
		if hashTable[runes[lowestIndexSingleChar]] == 1 {
			return runes[lowestIndexSingleChar]
		} else {
			lowestIndexSingleChar++
		}
	}

	return '?'
}
