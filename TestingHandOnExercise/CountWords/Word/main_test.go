package word

import (
	"testing"
)

func TestCount(t *testing.T) {
	x := Count("Ich funktioniere meaga gut")
	if x != 4 {
		t.Errorf("Expected: 4, got %v", x)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Ich bin ein etwas längerer String lol")
	}
}

func TestUseCount(t *testing.T) {
	x := UseCount("Ich bin ebenfalls ein längerer String. String länger.")
	_, ok := x["bin"]
	if !ok {
		t.Errorf("Expected true, got %v", ok)
	}
}

func BenchmarkUserCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Ich bin ebenfalls ein längerer String. String länger.")
	}
}
