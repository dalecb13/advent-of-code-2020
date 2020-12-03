# Advent of Code 2020

[Advent of Code 2020](https://adventofcode.com/2020) is the 5th iteration of an annual series of coding puzzles. This is my repository which hosts the solutions that I have come up with using Golang.

## Structure

This repository creates a package for each day and allows users to run the solver for a given day/problem via the command line. A package of helpers common to multiple problems can be found in the `helpers` folder.

## Usage

After cloning the repository, you should do the following:

1. Make sure you have Golang installed on your machine. Preferably a version which will allow you to use modules instead of vendors.
2. Run `go install` in each of the `solution/d*` folders. This will create a binary for each day which can then be run by the main routine.
3. Run `go install` and `go build` in the root folder of this project
4. Run `go run main.go ${d#p#}`, where ${d#p#} follows the following pattern
   1. `d1p1` for Day 1, Problem 1
   2. `d1p2` for Day 1, Problem 2
   3. `d2p1` for Day 2, Problem 1
   4. And so on
