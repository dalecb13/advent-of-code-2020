package d2

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

// validatePasswords takes a collection of strings.
// those strings can be split into two parts: Rules and Passwords
// Parse the Rule and Password from each line.
// Test the password with the rules to make sure that the passwords are valid
// Return the number of passwords which follow their rule
func ValidatePasswords(rulePasswordPairs []string) (int, error) {
	if len(rulePasswordPairs) < 1 {
		return -1, errors.New("Input does not have any rules or passwords")
	}

	numValid := 0

	// loop through password pairs
	for _, s := range rulePasswordPairs {
		// split the string into Rule & password
		rule, password, err := RuleSplit(s)
		if err != nil {
			return -1, err
		}

		isGood, err := isFollowRule(rule, password)

		if err != nil {
			return -1, err
		}

		if isGood {
			numValid++
		}
	}

	return numValid, nil
}

// Describes the giudelines for a password
// The ruleString looks as follows:
// lowerBound-upperbound character
type Rule struct {
	LowerBound   int
	UpperBound   int
	RequiredChar rune
}

type RulePair struct {
	Rule     Rule
	Password string
}

func RuleSplit(input string) (Rule, string, error) {
	if len(input) < 8 {
		return Rule{}, "", errors.New("Invalid input")
	}

	splits := strings.Split(input, ":")
	if len(splits) != 2 {
		log.Fatalln("Invalid number of colons")
	}

	ruleString := splits[0]
	password := splits[1]

	// parse ruleString into a Rule
	rule, err := parseRuleString(ruleString)
	if err != nil {
		return Rule{}, "", err
	} else {
		return rule, password, nil
	}
}

func parseRuleString(input string) (Rule, error) {
	if len(input) < 5 {
		return Rule{}, errors.New("Input string is too short to be parsed into a rule")
	}

	spaceSplit := strings.Split(input, " ")
	if len(spaceSplit) != 2 {
		return Rule{}, errors.New("Too many spaces in input")
	}

	requiredChar := spaceSplit[1]
	bounds := spaceSplit[0]

	hyphenSplit := strings.Split(bounds, "-")
	if len(hyphenSplit) != 2 {
		return Rule{}, errors.New("Invalid number of hyphens in Rule")
	}
	lowerBound := hyphenSplit[0]
	parsedLowerBound, eLower := strconv.Atoi(lowerBound)
	if eLower != nil {
		return Rule{}, errors.New("Problem with converting to string")
	}
	upperBound := hyphenSplit[1]
	parsedUpperBound, eUpper := strconv.Atoi(upperBound)
	if eUpper != nil {
		return Rule{}, errors.New("Problem with converting to string")
	}

	runes := []rune(requiredChar)

	return Rule{
		LowerBound:   parsedLowerBound,
		UpperBound:   parsedUpperBound,
		RequiredChar: runes[0],
	}, nil
}

// isFollowRule takes a given rule and password and checks to see
// if the password follows the rules
func isFollowRule(rule Rule, password string) (bool, error) {
	lowerBound := rule.LowerBound
	upperBound := rule.UpperBound
	char := rule.RequiredChar
	instances := 0

	for _, s := range password {
		if s == char {
			instances++
		}
	}

	return instances <= upperBound && instances >= lowerBound, nil
}
