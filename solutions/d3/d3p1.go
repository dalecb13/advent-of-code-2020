package d3

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/prometheus/common/log"
)

/*
	Toboggan Trajectory
	Input is a multi-line file with a certain pattern for each line.
	"." symbolizes an open spot.
	"#" symbolizes a tree

	Given a certain pattern of moving across the arrays, count the number of trees
*/

// D3p1 takes in the tree map and calculates the number of trees by going down 1 and right 3.
// The input `slopeMap` is assumed to be the first part of a pattern, so the first step is to duplicate the pattern.
func D3p1(slopeMap []string) (int, error) {
	if len(slopeMap) < 1 {
		return -1, errors.New("Not enough lines in input")
	}

	// The problem here is that when visiting a line,
	// we need to check if it is even possible to go any futher
	// It seems like a recursive traverse and check, where the exit case is that visiting element N causes an out-of-bound error
	// Can also pre-compute the number of slopes to expect if we assume that each slope is the same length

	// Pattern duplication
	//

	// start at top left
	lineIndex := 0
	elementIndex := 0
	isEnd := false
	numTree := 0

	for !isEnd {
		slope := slopeMap[lineIndex]
		fmt.Println("Checking slope " + strconv.Itoa(lineIndex) + ": " + slope)
		fmt.Printf("%q\n", slope)
		// fmt.Println("slope[" + strconv.Itoa(elementIndex) + "] = " + slope[elementIndex])
		check, err := IsTree(slopeMap[lineIndex], elementIndex)
		if err != nil {
			log.Warnln(err)
			isEnd = true
		} else if check {
			numTree++
			lineIndex++
			elementIndex += 3
		} else {
			lineIndex++
			elementIndex += 3
		}
	}

	return numTree, nil
}

// calcWcalcPatternReps calculates the number of times the slopePattern needs to be repeated
func CalcPatternReps(patternLength int, numSlopes int, numDown int, numRight int) int {
	// The relationship of the length of a slope can be determined from
	// * the number of slopes
	// * number of slopes to go down
	// * and number of spaces to go across
	// Basically the end result will be a rectangle.
	// The end rectangle will be proportional to the rectangle described by numDown and numRight.
	// The following proportion should be maintained:
	//             ${big_width}/${big_height} = ${small_width}/${small_height}
	// Where `small_width` = numRight & `small_height` = numDown
	// and `big_height` = numSlopes
	// then `big_width` = ${small_width} * ${big_height} / ${small_height}
	//
	// The algorithm also needs to consider the case where resulting `big_width` is not a whole number.
	// Perhaps another quirk with the problem will give us insight.
	//
	// As we are dealing with discrete units, the small rectangle will not tessellate nicely.
	// By visualizing the initial traversal, we can see that the bottom right corner will overlap
	//
	// With "3 across, 1 down" as an example, the tessellation should look like this
	// ___________________________________
	// |   .   #   .   . | .   .   .   .
	//                ________________
	// |   .   #   # | # | .   .   . | #
	// |_________________|
	// |   #   .   . | .   .   .   # | .
	// |             _________________
	//
	// From this insight, we can refine the algorithm which calculates the number of times to repeat the slopePattern.
	//
	// First we have the ratio of the `big_height` to the `small_height`.
	// Do we round down or up to the nearest multiple of `small_height`?
	// Well we do not want to increase the overall number of slopes (big_height),
	// so it should round down.
	//
	// From the above analysis, it should also be obvious that we are assuming that
	// numDown is less than numSlopes, and numRight is less than patternLength

	ratio := math.Floor(float64(numSlopes) / float64(numDown))

	// The vertical ratio should be used for the horizontal scaling as well.
	// Given the `small_width` we can calculate the end length.

	endLength := numRight * int(ratio)

	// From the end length, we can find the number of instances the slopePattern needs to be repeated.

	numPatternRepeats := int(float64(endLength) / float64(patternLength))

	return numPatternRepeats

}

// duplicatePattern duplicates each string a number of times.
// That number of times is determined by the
func DupePattern(slopes []string, numTimes int) ([]string, error) {
	dupedSlopes := []string{}

	for _, slope := range slopes {
		var builder strings.Builder
		for i := 0; i < numTimes; i++ {
			builder.WriteString(slope)
		}
		dupedSlopes = append(dupedSlopes, builder.String())
	}

	return dupedSlopes, nil
}

// isTree checks for three possibilities:
// * Tree
// * Toboggan-able terrain
// * Invalid element
// Comparing parts of a string in Golang is actually comparing bytes
// * byte representation of '.' is '\x2e'
// * byte representation of '#' is '\x23'
func IsTree(slope string, index int) (bool, error) {
	if len(slope) < index {
		errorMessage := "Out of bounds: (" + strconv.Itoa(len(slope)) + ") " + slope + " " + strconv.Itoa(index)
		return false, errors.New(errorMessage)
	} else {
		thing := slope[index]
		if thing == '.' {
			return false, nil
		} else if thing == '#' {
			return true, nil
		} else {
			errorMessage := "Unknown character at line: " + strconv.Itoa(index)
			return false, errors.New(errorMessage)
		}
	}
}
