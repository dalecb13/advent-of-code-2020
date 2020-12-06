package d6

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
	Customs Declaration form

	* There are 26 questions, a-z.
	* Each group of passengers only need to fill out one form.
	* Each question is a yes/no question.
	* Input from the file follows the corresponding rules:
		* Each group is separated by two newlines
		* Each person is separated by a single newline
	* Overlapping answers from different people within a group can be ignored.
	* Count the number of questions to which the group answered yes.
	* Count the number of questions to which the entire plane answered yes.
*/

func D6p1() {
	fileBytes, err := ioutil.ReadFile("data/day6.txt")
	if err != nil {
		log.Fatalln("Error reading in file: ", err)
	}

	groupData := strings.Split(string(fileBytes), "\n\n")

	numYes := 0
	for _, groupDatum := range groupData {
		// Use a map to keep track of a group's data.
		// A map will retain unique data and the keys act as a "set"
		groupMap := make(map[rune]bool)

		// groupAnswers is a []string, where each groupAnswer is the answer of a person
		groupAnswers := strings.Fields(groupDatum)

		for _, personAnswers := range groupAnswers {
			// the answer of a person is a string of characters--we put those in a map
			for _, personAnswer := range personAnswers {
				groupMap[personAnswer] = true
			}
		}

		// Add the group's number of "yes" answers to the result of the entire airplane
		numYes += len(groupMap)
	}

	log.Println(strconv.Itoa(numYes))
}
