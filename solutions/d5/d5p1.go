package d5

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var d5File = "data/day5.txt"

/*
	Boarding Pass Inference

	Lost my boarding pass, but first have to learn "binary space partitioning".
	"BSP" is the airline's way of determining seating locations.
	Seat locations are described by a string that is split into two parts:
	* First 7 characters are dedicated to row
	* Last 3 characters are dedicated to seat number within the row

	7 characters can represent the range 0-127
	* F means to take the front half
	* B means to take the back half

	3 characters can represent the range 0-7
	* L means to take the left half
	* R means to take the right half

	The goal for d5p1 is to determine the highest boarding pass in my input.
	when it suggests "highest", I am thinking that the seating code should have the most B's in the front,
	and the most R's in the front
*/

// FileToString reads in a file as a slice of strings
func FileToString(fp string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return []string{}, err
	}
	seatingCodes := strings.Split(string(fileBytes), "\n")
	return seatingCodes, nil
}

// CodeToSeat figures out the row and seat numbers for a given seating code
func CodeToSeat(code string) (int, int, error) {
	// split into 7 and 3
	rowCode := code[0:7]
	seatCode := code[7:10]

	rowNum, err := ParseRow(rowCode)
	if err != nil {
		return -1, -1, err
	}
	seatNum, err := ParseSeat(seatCode)
	if err != nil {
		return -1, -1, err
	}

	return rowNum, seatNum, nil
}

// Halfinate calculates the truncated value of the average of two numbers
func Halfinate(a int, b int) int {
	avg := (float64(b) + float64(a)) / 2.0
	intified := int(avg)
	return intified
}

// Given the binary space partitioning understanding, this function finds the exact row for a given code
func ParseRow(rowCode string) (int, error) {
	left := 0
	right := 127

	for _, code := range rowCode {
		if code == 'F' {
			// take left half by recalculating `right`
			right = Halfinate(left, right)
		} else if code == 'B' {
			// take right half by recalculating `left`
			left = Halfinate(left, right)
		}
	}

	return right, nil
}

// Given the binary space partitioning understanding, this function finds the exact seat for a given code
func ParseSeat(rowCodes string) (int, error) {
	// There's only 8 possibilities.
	left := 0
	right := 7

	for _, rowCode := range rowCodes {
		if rowCode == 'L' {
			// take left half by recalculating `right`
			right = Halfinate(left, right)
		} else if rowCode == 'R' {
			// take right half by recalculating `left`
			left = Halfinate(left, right)
		} else {
			return -1, errors.New(fmt.Sprintln("Unknown code: ", rowCode))
		}
	}

	return right, nil
}

// CalcSeatId calculates the seat ID by doing `rowNum` * 8 + `seatNum`
func CalcSeatId(rowNum int, seatNum int) int {
	if rowNum < 0 || seatNum < 0 {
		return -1
	} else {
		return rowNum*8 + seatNum
	}
}

func D5p1() {
	codes, err := FileToString(d5File)
	if err != nil {
		log.Fatalln(err)
	}

	ids := []int{}

	for _, code := range codes {
		rowNum, seatNum, err := CodeToSeat(code)
		if err != nil {
			log.Fatalln(err)
		}
		if rowNum == -1 {
			log.Fatalln("Error parsing row number: ", rowNum)
		}
		if seatNum == -1 {
			log.Fatalln("Error parsing seat number: ", seatNum)
		}

		id := CalcSeatId(rowNum, seatNum)
		ids = append(ids, id)
	} // end for loop

	// sort to figure out highest ID value
	sort.Ints(ids)

	log.Println(strconv.Itoa(ids[len(ids)-1]))
}
