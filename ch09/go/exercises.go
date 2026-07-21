package ch09

func ReverseString(str string) string {
	s := &Stack{}

	for _, char := range str {
		s.Push(char)
	}

	reversedChars := []rune{}

	for {
		if s.Read() != nil {
			reversedChars = append(reversedChars, s.Pop().(rune))
		} else {
			break
		}
	}

	return string(reversedChars)
}
