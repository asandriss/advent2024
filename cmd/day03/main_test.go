package main

import "testing"

func TestSolve(t *testing.T) {
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	expect := 48
	actual := solve(input)

	if actual != expect {
		t.Errorf("Solve(input) = %d, want %d", actual, expect)
	}
}
