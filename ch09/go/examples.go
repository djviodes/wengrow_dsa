package ch09

import (
	"fmt"
	"slices"
)

type Stack struct {
	data []any
}

func (s *Stack) Push(element any) {
	s.data = append(s.data, element)
}

func (s *Stack) Pop() any {
	var element any

	if len(s.data) > 0 {
		element = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
	}

	return element
}

func (s *Stack) Read() any {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	} else {
		return nil
	}
}

type Linter struct {
	stack Stack
}

func (l *Linter) Lint(text string) string {
	for _, char := range text {
		if l.isOpeningBrace(char) {
			l.stack.Push(char)
		} else if l.isClosingBrace(char) {
			poppedOpeningBrace := l.stack.Pop()

			if poppedOpeningBrace == nil {
				return string(char) + " doesn't have an opening brace"
			}

			if l.isNotMatch(poppedOpeningBrace.(rune), char) {
				return string(char) + " has mismatched opening brace"
			}
		}
	}

	if l.stack.Read() != nil {
		return string(l.stack.Read().(rune)) + " does not have a closing brace"
	}

	return "No errors"
}

func (l *Linter) isOpeningBrace(char rune) bool {
	openingBraces := []rune{'(', '[', '{'}

	return slices.Contains(openingBraces, char)
}

func (l *Linter) isClosingBrace(char rune) bool {
	closingBraces := []rune{')', ']', '}'}

	return slices.Contains(closingBraces, char)
}

func (l *Linter) isNotMatch(openingBrace rune, closingBrace rune) bool {
	if (openingBrace == '(' && closingBrace == ')') ||
		(openingBrace == '[' && closingBrace == ']') ||
		(openingBrace == '{' && closingBrace == '}') {
		return false
	}

	return true
}

type Queue struct {
	data []any
}

func (q *Queue) Enqueue(element any) {
	q.data = append(q.data, element)
}

func (q *Queue) Dequeue() any {
	var element any

	if len(q.data) > 0 {
		element = q.data[0]
		q.data = q.data[1:]
	}

	return element
}

func (q *Queue) Read() any {
	if len(q.data) > 0 {
		return q.data[0]
	} else {
		return nil
	}
}

type PrintManager struct {
	queue Queue
}

func (pm *PrintManager) QueuePrintJob(document any) {
	pm.queue.Enqueue(document)
}

func (pm *PrintManager) Run() {
	for {
		if pm.queue.Read() != nil {
			pm.print(pm.queue.Dequeue())
		} else {
			break
		}
	}
}

func (pm *PrintManager) print(document any) {
	fmt.Println("Document:", document)
}
