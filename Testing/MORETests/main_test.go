package main

import "testing"

func TestMultiply(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{-1, 1, 2}, -2},
		{[]int{2, 7}, 14},
		{[]int{-10, -10}, 100},
		{[]int{4, 8}, 32},
		{[]int{10, 10}, 100},
		{[]int{10, 10, 10}, 1000},
		{[]int{0, 0, 0}, 0},
	}

	for _, v := range tests {
		x := Multiply(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
