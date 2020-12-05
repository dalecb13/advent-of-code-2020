package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dalecb13/aoc2020/d1"
	"github.com/dalecb13/aoc2020/d2"
	"github.com/dalecb13/aoc2020/d3"
	"github.com/dalecb13/aoc2020/d4"
	"github.com/dalecb13/aoc2020/helpers"
)

var d1File = "data/day1.txt"
var d2File = "data/day2.txt"
var d3File = "data/day3.txt"

func main() {
	log.Println("The CLI for testing solutions to the Advent of Code 2020!")

	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	problem := os.Args[1]

	switch problem {
	case "d1p1":
		fmt.Println("Day 1, Problem 1")
		// Parse file input
		expenseReport, err := helpers.FileOfInts(d1File)
		if err != nil {
			log.Fatalln("Issue reading file", err)
		}
		// Calculate solution
		i, j, err := d1.D1p1(expenseReport)
		if err != nil {
			log.Fatalln("No pair found")
		}
		log.Println("D1P1 result: ", strconv.Itoa(i*j))
	case "d1p2":
		fmt.Println("Day 1, Problem 2")
		// Parse file input
		expenseReport, err := helpers.FileOfInts(d1File)
		if err != nil {
			log.Fatalln("Issue reading file", err)
		}
		// Calculate solution
		one, two, three, e := d1.D1p2(expenseReport)
		if e != nil {
			log.Fatalln("No trio found")
		}
		log.Println("D1P2 result: ", strconv.Itoa(one*two*three))
	case "d2p1":
		log.Println("Day 2, Problem 1")
		// Parse file input
		passwords, err := helpers.FileOfStrings(d2File)
		if err != nil {
			log.Fatalln("Error reading file", err)
		}
		// Calculate solution
		numValidPasswords, err := d2.ValidatePasswords(passwords)
		if err != nil {
			log.Fatalln("Error parsing passwords", err)
		}
		log.Println("D2P1 result: ", strconv.Itoa(numValidPasswords))
	case "d2p2":
		log.Println("Day 2, Problem 2")
		// Parse file input
		passwords, err := helpers.FileOfStrings(d2File)
		if err != nil {
			log.Fatalln("Error reading file", err)
		}
		// Calculate solution
		numValidPasswords, err := d2.ValidateDay2Passwords(passwords)
		if err != nil {
			log.Fatalln("Error parsing passwords", err)
		}
		log.Println("D2P2 result: ", strconv.Itoa(numValidPasswords))
	case "d3p1":
		log.Println("Day 3, Problem 1")
		// Parse file input
		slopeMap, err := helpers.FileOfStrings(d3File)
		if err != nil {
			log.Fatalln("Error reading file", err)
		}
		// Calculate solution
		numTrees, err := d3.D3p1simple(slopeMap, 1, 3)
		if err != nil {
			log.Fatalln("Error running d3p1 problem ", err)
		} else {
			log.Println(strconv.Itoa(numTrees))
		}
	case "d3p2":
		log.Println("Day 3, Problem 2")
		// Parse file input
		slopeMap, err := helpers.FileOfStrings(d3File)
		if err != nil {
			log.Fatalln("Error reading file", err)
		}
		// Calculate solution
		product, err := d3.D3p2(slopeMap)
		if err != nil {
			log.Fatalln("Error running d3p2 problem ", err)
		} else {
			log.Println(strconv.Itoa(product))
		}
	case "d4p1":
		d4.D4p1()
	case "d4p2":
		d4.D4p2()
	default:
		log.Fatalln("Unknown argument found: ", problem)
	}
}
