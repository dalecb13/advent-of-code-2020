package d3

import (
	"errors"
	"strconv"
)

type TraversePattern struct {
	numDown  int
	numRight int
}

// D3p2 is similar to D3p1 except that it runs D3p1 multiple times with different traversing patterns.
// The end solution is the product of the results of D3p1
func D3p2(slopeMap []string) (int, error) {
	patterns := []TraversePattern{
		{
			numDown:  1,
			numRight: 1,
		},
		{
			numDown:  1,
			numRight: 3,
		},
		{
			numDown:  1,
			numRight: 5,
		},
		{
			numDown:  1,
			numRight: 7,
		},
		{
			numDown:  2,
			numRight: 1,
		},
	}
	product := 1

	for i, pattern := range patterns {
		numTrees, err := D3p1simple(slopeMap, pattern.numDown, pattern.numRight)
		if err != nil {
			return -1, errors.New("Issue with traversing with pattern[" + strconv.Itoa(i) + "]")
		} else {
			product *= numTrees
		}
	}

	return product, nil
}
