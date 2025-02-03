package mymath

import (
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		ehm []int
		answer float64
	}

	tests := []test{
		{[]int{1, 4, 6, 8}, 5},
		{[]int{10, 1, 10, 8}, 9},
		{[]int{1, 10, 12}, 10},
		{[]int{5,10, 15, 20, 30}, 15},
	}

	for _, v := range tests {
		x := CenteredAvg(v.ehm)
		if x != v.answer {
			t.Errorf("Expected: %v, got %v", v.answer, x)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	CenteredAvg([]int{10, 10, 4, 1, 5})
}
