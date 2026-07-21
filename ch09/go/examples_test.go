package ch09

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}

	if result := s.Pop(); result != nil {
		t.Errorf("expected Pop on empty stack to return nil, got %v", result)
	}

	if result := s.Read(); result != nil {
		t.Errorf("expected Read on empty stack to return nil, got %v", result)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if result := s.Read(); result != 3 {
		t.Errorf("expected Read to return 3, got %v", result)
	}

	if result := s.Pop(); result != 3 {
		t.Errorf("expected Pop to return 3, got %v", result)
	}

	if result := s.Read(); result != 2 {
		t.Errorf("expected Read to return 2, got %v", result)
	}
}

func TestLinter(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"(a)", "No errors"},
		{"(a[b]{c})", "No errors"},
		{"(a", "( does not have a closing brace"},
		{"(a]", "] has mismatched opening brace"},
		{"a)", ") doesn't have an opening brace"},
	}

	for _, c := range cases {
		l := &Linter{}
		if result := l.Lint(c.input); result != c.want {
			t.Errorf("Lint(%q) = %q, want %q", c.input, result, c.want)
		}
	}
}

func TestQueue(t *testing.T) {
	q := &Queue{}

	if result := q.Dequeue(); result != nil {
		t.Errorf("expected Dequeue on empty queue to return nil, got %v", result)
	}

	if result := q.Read(); result != nil {
		t.Errorf("expected Read on empty queue to return nil, got %v", result)
	}

	q.Enqueue("first.pdf")
	q.Enqueue("second.pdf")
	q.Enqueue("third.pdf")

	if result := q.Read(); result != "first.pdf" {
		t.Errorf("expected Read to return \"first.pdf\", got %v", result)
	}

	if result := q.Dequeue(); result != "first.pdf" {
		t.Errorf("expected Dequeue to return \"first.pdf\", got %v", result)
	}

	if result := q.Read(); result != "second.pdf" {
		t.Errorf("expected Read to return \"second.pdf\", got %v", result)
	}
}

func ExamplePrintManager_Run() {
	pm := &PrintManager{}
	pm.QueuePrintJob("first.pdf")
	pm.QueuePrintJob("second.pdf")
	pm.QueuePrintJob("third.pdf")
	pm.Run()
	// Output: Document: first.pdf
	// Document: second.pdf
	// Document: third.pdf
}
