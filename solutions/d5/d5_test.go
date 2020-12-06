package d5

import (
	"strconv"
	"testing"
)

func TestHalfinate(t *testing.T) {
	left := 1
	right := 3

	actualResult := Halfinate(left, right)

	if actualResult != 2 {
		t.Errorf("Error calculating Halfinate: %d", actualResult)
	}
}

func TestParseSeat(t *testing.T) {
	codes := []string{"LLL", "LLR", "LRL", "LRR", "RLL", "RLR", "RRL", "RRR"}
	expectedResults := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for i, code := range codes {
		actualResult, err := ParseSeat(code)
		if err != nil {
			t.Errorf("Error parsing code: %s", err)
		} else if actualResult != expectedResults[i] {
			t.Errorf("Actual result %s does not match expected result %s", strconv.Itoa(actualResult), strconv.Itoa(expectedResults[i]))
		}
	}
}
