package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dalecb13/aoc2020/d1"
	"github.com/dalecb13/aoc2020/helpers"
)

var d1File = "data/day1.txt"

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
		one, two, three, e := d1.D1p2(expenseReport)
		if e != nil {
			log.Fatalln("No trio found")
		}
		log.Println("D1P2 result: ", strconv.Itoa(one*two*three))
	}
}
