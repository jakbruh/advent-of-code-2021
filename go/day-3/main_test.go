package main

import "testing"

func TestPartOne(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	total := partOne(input)
	if total != 198 {
		t.Errorf("Ouput was incorrect, got: %d, want: %d.", total, 198)
	}
}
