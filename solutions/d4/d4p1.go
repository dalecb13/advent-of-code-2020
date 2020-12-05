package d4

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
	Passport verification

	Passports must have the required fields:
	* byr (Birth Year)
	* iyr (Issue Year)
	* eyr (Expiration Year)
	* hgt (height)
	* hcl (hair color)
	* ecl (eye color)
	* pid (passport ID)
	* cid (country ID)

	Properties are separated by spaces.
	Passports are separated by two newlines.

	Valid passports:
	* Must have all fields present (cid is the only optional field)

	https://adventofcode.com/2020/day/4
*/

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportId     string
	CountryId      string
}

func PropsToValidPassport(properties []string) ([]string, error) {
	if len(properties) < 7 {
		return []string{}, nil
	}

	m := make(map[string]string)
	for _, property := range properties {
		split := strings.Split(property, ":")
		key := split[0]
		val := split[1]

		m[key] = val
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	numKeys := len(keys)

	if numKeys == 8 {
		// is a passport has 8 keys, it is valid
		return properties, nil
	} else if numKeys == 7 {
		// if a Passport has 7 keys, it could be valid
		// check to see that country id is the only value missing
		value := m["cid"]
		if value == "" {
			// if cid is missing, then it's a good passport
			return properties, nil
		} else {
			// if cid is not missing, then it's not a good passport
			return []string{}, nil
		}
	} else {
		// anything less than 7 keys is automatically invalid
		return []string{}, nil
	}
}

func PropsToPassport(properties []string) (Passport, error) {
	var passport Passport
	for _, property := range properties {
		kv := strings.Split(property, ":")
		key := kv[0]
		value := kv[1]

		switch key {
		case "byr":
			byr, err := strconv.Atoi(value)
			if err != nil {
				return Passport{}, err
			}
			passport.BirthYear = byr
		case "iyr":
			iyr, err := strconv.Atoi(value)
			if err != nil {
				return Passport{}, err
			}
			passport.IssueYear = iyr
		case "eyr":
			eyr, err := strconv.Atoi(value)
			if err != nil {
				return Passport{}, err
			}
			passport.ExpirationYear = eyr
		case "hgt":
			passport.Height = value
		case "hcl":
			passport.HairColor = value
		case "ecl":
			passport.EyeColor = value
		case "pid":
			passport.PassportId = value
		case "cid":
			passport.CountryId = value
		default:
			log.Fatalln("Unknown property: ", key)
		} // end switch
	} // end for loop

	return passport, nil
}

func ParsePassportStrings(fp string) ([][]string, error) {
	fileBytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return [][]string{}, err
	}
	passportStrings := strings.Split(string(fileBytes), "\n\n")

	var validPassports [][]string

	for _, passportString := range passportStrings {
		properties := strings.Fields(passportString)
		validProps, err := PropsToValidPassport(properties)
		if err != nil {
			return [][]string{}, err
		} else if len(validProps) != 0 {
			validPassports = append(validPassports, validProps)
		}
	} // end for loop

	return validPassports, nil
}

func ParsePassports(fp string) ([]Passport, error) {
	fileBytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return []Passport{}, err
	}
	passportStrings := strings.Split(string(fileBytes), "\n\n")

	passports := []Passport{}

	for _, passportString := range passportStrings {
		properties := strings.Fields(passportString)

		passport, err := PropsToPassport(properties)
		if err != nil {
			return []Passport{}, err
		} else {
			passports = append(passports, passport)
		}
	} // end for loop

	return passports, nil
}

// ValidatePassport validates the provided passport.
// According to the problem description, all fields are mandatory except for CountryId.
// According to Golang,
func IsPassportValid(passport Passport) (bool, error) {
	if passport.BirthYear != 0 && passport.IssueYear != 0 && passport.ExpirationYear != 0 && passport.EyeColor != "" && passport.HairColor != "" && passport.Height != "" && passport.PassportId != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func ValidatePassports(passports []Passport) ([]Passport, error) {
	if len(passports) < 1 {
		return []Passport{}, errors.New("Input is undefined")
	}

	var validPassports []Passport

	for _, passport := range passports {
		isValid, err := IsPassportValid(passport)
		if err != nil {
			return []Passport{}, err
		} else if isValid {
			validPassports = append(validPassports)
		}
	} // end for loop

	return validPassports, nil
}

func D4p1() {
	// Read in file and parse as passports
	passports, err := ParsePassportStrings("data/day4.txt")
	if err != nil {
		log.Fatalln("Error parsing file: ", err)
	}

	numValid := len(passports)
	log.Println(strconv.Itoa(numValid))
}
