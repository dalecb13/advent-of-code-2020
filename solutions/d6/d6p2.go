package d6

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
	Similar to Day 6 Problem 1, but the question only counts for those that everyone answered yes.

	Instead of a map[rune]bool, can use a map[rune]int where the value stores the number of people voting for that question.
*/

func D6p2() {
	fileBytes, err := ioutil.ReadFile("data/day6.txt")
	if err != nil {
		log.Fatalln("Error reading in file: ", err)
	}

	groupData := strings.Split(string(fileBytes), "\n\n")

	numYes := 0
	for _, groupDatum := range groupData {
		// Use a map to keep track of a group's data.
		// A map will retain unique data and the keys act as a "set"
		groupMap := make(map[rune]int)

		// groupAnswers is a []string, where each groupAnswer is the answer of a person
		groupAnswers := strings.Fields(groupDatum)
		numInGroup := len(groupAnswers)

		for _, personAnswers := range groupAnswers {
			// the answer of a person is a string of characters--we put those in a map
			for _, personAnswer := range personAnswers {
				groupMap[personAnswer]++
			}
		}

		// Find the entries which have the same value as the number of people voting for it
		for _, v := range groupMap {
			if v == numInGroup {
				numYes++
			}
		}
	}

	log.Println(strconv.Itoa(numYes))
}
