package main

import "testing"

func TestSubtract(t *testing.T) {
	total := subtract(5, 5)
	if total != 0 {
		t.Errorf("Sum was incorrect, got %d, want %v", total, 0)
	}
}

func TestDoMath(t *testing.T) {
	total := doMath(4, 4, add)
	if total != 8 {
		t.Errorf("Sum was incorrect, got %d, want %v", total, 8)
	}
	total = doMath(4, 4, subtract)
	if total != 0 {
		t.Errorf("Sum was incorrect, got %d, want %v", total, 0)
	}
}
