package helpers

import (
	"io/ioutil"
	"strconv"
	"strings"
)

/*
	Collection of helper functions
*/

func FileOfStrings(fp string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fp)
	if err != nil {
		return []string{}, err
	}
	sliceData := strings.Split(string(fileBytes), "\n")
	return sliceData, nil
}

func FileOfInts(fp string) ([]int, error) {
	ss, err := FileOfStrings(fp)
	if err != nil {
		return []int{}, err
	}

	var is = []int{}
	for _, s := range ss {
		converted, e := strconv.Atoi(s)
		if e != nil {
			return []int{}, e
		}
		is = append(is, converted)
	}

	return is, nil
}
