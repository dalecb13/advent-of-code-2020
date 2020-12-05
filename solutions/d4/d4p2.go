package d4

import (
	"log"
	"regexp"
	"strconv"
)

/*
	This problem requires validations for all of the fields.
	Country ID is still optional.

	The validations include:
	* 2002 >= Birth Year >= 1920
	* 2020 >= Issue Year >= 2010
	* 2030 >= Expiration Year >= 2020
	* Height
		* Must have "in" or "cm" units
		* if "cm", 193 >= height >= 150
		* if "in", 76 >= height >= 59
	* Hair Color: must be a hexadecimal value (with pound sign)
	* Eye Color: amb blu brn gry grn hzl oth
	* Passport ID: Nine digit number, leading zeros okay
	* Country ID: still optional
*/

func IsValidBirthYear(year int) bool {
	return 2002 >= year && year >= 1920
}

func IsValidIssueYear(year int) bool {
	return 2020 >= year && year >= 2010
}

func IsValidExpirationYear(year int) bool {
	return 2030 >= year && year >= 2020
}

func IsValidHeight(height string) bool {
	bs := []byte(height)
	l := len(bs)

	// height string must be at least 4 units long
	if len(height) < 4 {
		return false
	}

	// get number portion
	numBs := bs[0 : l-2]
	numStr := string(numBs)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return false
	}
	// get last two values of the byte slice
	dim := bs[l-2 : l]
	dimStr := string(dim)

	// last two values of byte slice must be "cm" or "in"
	if dimStr == "cm" {
		return 193 >= num && num >= 150
	} else if dimStr == "in" {
		return 76 >= num && num >= 59
	} else {
		return false
	}
}

func IsValidHairColor(hairColor string) bool {
	if len(hairColor) != 7 {
		return false
	}

	if hairColor[0] != '#' {
		return false
	}

	bs := []byte(hairColor)
	postPound := bs[1:7]

	matched, err := regexp.Match(`[0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F]`, postPound)
	if err != nil && !matched {
		return false
	} else {
		return true
	}

}

func IsValidEyeColor(eyeColor string) bool {
	return eyeColor == "amb" || eyeColor == "blu" || eyeColor == "brn" || eyeColor == "gry" || eyeColor == "grn" || eyeColor == "hzl" || eyeColor == "oth"
}

// IsValidPassportId checks to see if in the input consists of 9 numbers, including leading zeros
func IsValidPassportId(pid string) bool {
	converted, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return len(pid) == 9 && converted != 0
}

func StrictValidate(passportsProps [][]string) ([]Passport, error) {
	var passports []Passport
	for _, passportProps := range passportsProps {
		passport, err := PropsToPassport(passportProps)
		if err != nil {
			log.Fatalln("Error parsing properties to passport: ", err)
		}

		// run through validation
		if IsValidBirthYear(passport.BirthYear) && IsValidExpirationYear(passport.ExpirationYear) && IsValidEyeColor(passport.EyeColor) && IsValidHairColor(passport.HairColor) && IsValidHeight(passport.Height) && IsValidIssueYear(passport.IssueYear) && IsValidPassportId(passport.PassportId) {
			passports = append(passports, passport)
		}

	} // end for loop

	return passports, nil
}

func D4p2() {
	// Read in file and parse as passports
	passportsProps, err := ParsePassportStrings("data/day4.txt")
	if err != nil {
		log.Fatalln("Error parsing file: ", err)
	}

	passports, err := StrictValidate(passportsProps)
	if err != nil {
		log.Fatalln("Error validating passport properties: ", err)
	}

	log.Println(strconv.Itoa(len(passports)))
}
