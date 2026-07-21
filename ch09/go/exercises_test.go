package ch09

import "testing"

func TestReverseString(t *testing.T) {
	if result := ReverseString("hello"); result != "olleh" {
		t.Errorf("expected \"olleh\", got %q", result)
	}

	if result := ReverseString(""); result != "" {
		t.Errorf("expected \"\", got %q", result)
	}

	if result := ReverseString("racecar"); result != "racecar" {
		t.Errorf("expected \"racecar\", got %q", result)
	}

	// multi-byte character -- regression check for byte-vs-rune handling
	if result := ReverseString("André"); result != "érdnA" {
		t.Errorf("expected \"érdnA\", got %q", result)
	}
}
