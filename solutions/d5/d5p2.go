package d5

import (
	"log"
	"strconv"
)

// D5p2 figures out which seat is missing--that is my seat
// Each row as 8 seats, so it would make sense to initialize a map.
// The map would have a key being the row and a value being the seats taken in that row
// After putting everything in the map, figure out what's missing
func D5p2() {
	codes, err := FileToString(d5File)
	if err != nil {
		log.Fatalln(err)
	}

	// Map<int, Map<int, Map<int, bool>>
	// outerMap is a map whose values are maps
	// The values are maps because we want to make a collection that stores unique values
	outerMap := make(map[int]map[int]bool)

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

		if len(outerMap[rowNum]) == 0 {
			// row hasn't been added yet, so we add it
			newRow := make(map[int]bool)
			newRow[seatNum] = true
			outerMap[rowNum] = newRow
		} else {
			// row has been added, so we add the seatNum to the value
			existingRow := outerMap[rowNum]
			existingRow[seatNum] = true
			outerMap[rowNum] = existingRow
		}

	} // end for loop

	// traverse through map and find values which have less than 8 keys
	for row, seats := range outerMap {
		if len(seats) == 7 {
			// The row that has 7 seats is the row that we're looking for.
			// We just have to figure out which seat is empty.
			// Fortunately the inner map, the map of occupied seats, is sorted
			seatNumber := 0
			for index, _ := range seats {
				if index != seatNumber {
					// if seatNumber does not equal index, then we've found the empty seat
					// print out the Seat ID
					log.Println(strconv.Itoa(row*8 + seatNumber))
				}
			}
		}
	}
}
