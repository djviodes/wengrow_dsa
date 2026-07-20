package ch08

import "testing"

func TestListIntersection(t *testing.T) {
	result := ListIntersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8})
	want := []int{2, 4}

	if len(result) != len(want) {
		t.Fatalf("expected %v, got %v", want, result)
	}
	for i := range want {
		if result[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, result)
		}
	}
}

func TestDuplicateChars(t *testing.T) {
	result := DuplicateChars([]string{"a", "b", "c", "d", "c", "e", "f"})
	if result != "c" {
		t.Errorf("expected \"c\", got %q", result)
	}
}

func TestMissingLetters(t *testing.T) {
	result := MissingLetters("the quick brown box jumps over the lazy dog")
	if result != 'f' {
		t.Errorf("expected 'f', got %q", result)
	}
}

func TestNonDuplicatedString(t *testing.T) {
	if result := NonDuplicatedString("minimum"); result != 'n' {
		t.Errorf("expected 'n', got %q", result)
	}

	// multi-byte character before the answer -- regression check for the
	// byte-vs-rune indexing bug this function originally had
	if result := NonDuplicatedString("aaébb"); result != 'é' {
		t.Errorf("expected 'é', got %q", result)
	}

	if result := NonDuplicatedString("aabbcc"); result != '?' {
		t.Errorf("expected '?', got %q", result)
	}
}
