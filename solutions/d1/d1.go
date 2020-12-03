package d1

import (
	"errors"
)

// d1p1 is the function which solves the first puzzle on the first day of the Advent of Code for 2020.
// d1p1 acceps an array of numbers, finds the numbers which sum to 2020, and returns their product
// The procedure is an n^2 operation which loops through the array a total of "n" number of times to test if the sum of array[n] and array[m] equals the targeted number
func D1p1(expenseReport []int) (int, int, error) {
	if len(expenseReport) < 2 {
		return 0, 0, errors.New("Invalid size")
	}

	// Use a double for loop and brute-force combinations until one is found
	for _, s := range expenseReport {
		for _, t := range expenseReport {
			if s+t == 2020 {
				return s, t, nil
			}
		}
	}

	// else nothing found
	return -1, -1, errors.New("No pair found")
}

func D1p2(expenseReport []int) (int, int, int, error) {
	if len(expenseReport) < 2 {
		return -1, -1, -1, errors.New("Invalid size")
	}

	// Use a double for loop and brute-force combinations until one is found
	for _, s := range expenseReport {
		for _, t := range expenseReport {
			for _, u := range expenseReport {
				if s+t+u == 2020 {
					return s, t, u, nil
				}
			}
		}
	}

	// else nothing found
	return -1, -1, -1, errors.New("No pair found")
}
