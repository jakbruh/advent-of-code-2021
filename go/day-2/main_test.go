package main

import "testing"

func TestPartTwo(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	total := partTwo(input)
	if total != 900 {
		t.Errorf("Ouput was incorrect, got: %d, want: %d.", total, 900)
	}
}
