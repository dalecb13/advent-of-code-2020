package d4

import (
	"strconv"
	"testing"
)

func TestLeadingZeros(t *testing.T) {
	input := "0011223344"
	converted, err := strconv.Atoi(input)

	if err != nil || converted != 11223344 {
		t.Errorf("Conversion error")
	}
}
