package dog

import "testing"

func TestYears(t *testing.T) {
	type testAge struct {
		age    int
		answer int
	}

	tests := []testAge{
		{1, 7},
		{5, 35},
		{10, 70},
		{13, 91},
		{15, 105},
		{8, 56},
	}
	for _, v := range tests {
		x := Years(v.age)
		if x != v.answer {
			t.Errorf("Expected: %v, got: %v", v.answer, x)
		}

	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}
