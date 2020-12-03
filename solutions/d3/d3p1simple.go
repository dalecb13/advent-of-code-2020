package d3

import (
	"errors"
)

/*
	The solution here is to use the modulus operation to
	loop through the slope when counting trees.
*/

// D3p1simple takes in a slopeMap and figures out the number of trees
// when traversing in a certain direction
func D3p1simple(slopeMap []string, numDown int, numRight int) (int, error) {
	if len(slopeMap) < 1 {
		return -1, errors.New("Not enough lines in input")
	}
	if numDown < 1 || numRight < 0 {
		return -1, errors.New("Invalid traversing pattern")
	}

	lineIndex := 0
	elementIndex := 0
	numTree := 0
	isTraversing := true

	numSlopes := len(slopeMap)
	slopeLength := len(slopeMap[0])

	for isTraversing {
		// leaves the loop
		if lineIndex >= numSlopes {
			isTraversing = false
		} else {
			slope := slopeMap[lineIndex]
			idx := elementIndex % slopeLength
			terrain := slope[idx]

			if terrain == '#' {
				numTree++
			}

			lineIndex += numDown
			elementIndex += numRight
		}
	} // end for loop

	return numTree, nil
}
