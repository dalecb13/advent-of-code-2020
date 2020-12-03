package d3

import (
	"log"
	"strings"
	"testing"
)

func TestIsTree(t *testing.T) {
	// Setup test
	slope := ".##...#.#.......#.######....#."
	indices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	expectedResults := []bool{false, true, true, false, false, false, true, false, true, false, false, false, false, false, false, false, true, false, true, true, true, true, true, true, false, false, false, false, true, false}

	// Run table test
	for i, index := range indices {
		actualResult, err := IsTree(slope, index)
		if err != nil {
			log.Fatalln("Error with .IsTree()", err)
		}

		// Evaluate results
		if actualResult != expectedResults[i] {
			t.Errorf("For %d, expected %t, got %t\n", i, expectedResults[i], actualResult)
		}
	}
}

func TestCalcPatternReps(t *testing.T) {

}

func TestDupePattern(t *testing.T) {
	// Setup test
	slopes := []string{".#.", "#.#", "###", "..."}
	numTimes := 3
	expectedResults := []string{".#..#..#.", "#.##.##.#", "#########", "........."}

	// Run table test
	for i, slope := range slopes {
		var builder strings.Builder
		for i := 0; i < numTimes; i++ {
			builder.WriteString(slope)
		}

		// Evaluate results
		if builder.String() != expectedResults[i] {
			t.Errorf("For %d, expected %s, got %s", i, expectedResults[i], builder.String())
		}
	}
}
