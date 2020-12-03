package d2

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

/*
	Day 2 Problem 2 has a different set of rules.
	Instead of `n-m` being the lower/upper bound for the number of times
	the character appears, the new rule states that the given character
	has to appear ONLY ONCE in either the left or right index
*/

type Rule2 struct {
	index1 int
	index2 int
	char   rune
}

func Rule2Split(input string) (Rule2, string, error) {
	if len(input) < 8 {
		return Rule2{}, "", errors.New("Invalid input")
	}

	splits := strings.Split(input, ":")
	if len(splits) != 2 {
		log.Fatalln("Invalid number of colons")
	}

	ruleString := splits[0]
	password := splits[1]

	// parse ruleString into a Rule
	rule, err := parseRule2String(ruleString)
	if err != nil {
		return Rule2{}, "", err
	} else {
		return rule, password, nil
	}
}

// parseRule2String creates an object which describes the rule for Day 2 Problem 2
func parseRule2String(ruleString string) (Rule2, error) {
	if len(ruleString) < 5 {
		return Rule2{}, errors.New("Input string is too short to be parsed into a rule")
	}

	spaceSplit := strings.Split(ruleString, " ")
	if len(spaceSplit) != 2 {
		return Rule2{}, errors.New("Too many spaces in input")
	}

	requiredChar := spaceSplit[1]
	indices := spaceSplit[0]

	hyphenSplit := strings.Split(indices, "-")
	if len(hyphenSplit) != 2 {
		return Rule2{}, errors.New("Invalid number of hyphens in Rule")
	}
	lowerIndex := hyphenSplit[0]
	parsedLowerIndex, eLower := strconv.Atoi(lowerIndex)
	if eLower != nil {
		return Rule2{}, errors.New("Problem with converting to string")
	}
	upperIndex := hyphenSplit[1]
	parsedUpperIndex, eUpper := strconv.Atoi(upperIndex)
	if eUpper != nil {
		return Rule2{}, errors.New("Problem with converting to string")
	}

	runes := []rune(requiredChar)

	return Rule2{
		index1: parsedLowerIndex,
		index2: parsedUpperIndex,
		char:   runes[0],
	}, nil
}

// isFollowRule2 takes a given rule and password and checks to see
// if the password follows the rules
func isFollowRule2(rule Rule2, password string) (bool, error) {
	lowerIndex := rule.index1
	upperIndex := rule.index2
	char := rule.char

	runes := []rune(password)

	if (runes[lowerIndex] == char && runes[upperIndex] != char) || (runes[lowerIndex] != char && runes[upperIndex] == char) {
		return true, nil
	}

	return false, nil
}

func ValidateDay2Passwords(lines []string) (int, error) {
	if len(lines) < 1 {
		return -1, errors.New("Not enough lines in input")
	}

	numCorrect := 0

	for _, line := range lines {
		rule, password, err := Rule2Split(line)
		if err != nil {
			return -1, err
		}

		good, err := isFollowRule2(rule, password)
		if err != nil {
			return -1, err
		}

		if good {
			numCorrect++
		}
	}

	return numCorrect, nil
}
