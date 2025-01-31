package main

import "testing"

func TestSubtract(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{-1, 1, 2}, 2},
		{[]int{2, 7}, 9},
		{[]int{-10, -10}, -20},
		{[]int{4, 8}, 12},
		{[]int{10, 10}, 20},
		{[]int{10, 10, 10}, 30},
		{[]int{0, 0, 0}, 0},
	}

	for _, v := range tests {
		x := Subtract(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
