package ch10

import "fmt"

func ArrayParser(array []any) {
	if len(array) == 0 {
		return
	}

	if num, ok := array[0].(int); ok {
		fmt.Println(num)
		ArrayParser(array[1:])
	} else {
		ArrayParser(array[0].([]any))
		ArrayParser(array[1:])
	}
}
