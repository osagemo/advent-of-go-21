package puzzles

import "fmt"

type DayTen struct {
	day
}

type tagPair struct {
	opening rune
	closing rune
}

func (DayTen) GetPuzzleName() string {
	return "Day 10: Syntax Scoring"
}

func (d DayTen) PartOne() string {
	// validate lines for corruption
	return "not implemented"
}

func (d DayTen) validateLine(line string) {
	acceptedPairs := []tagPair{
		{'{', '}'},
		{'(', ')'},
		{'[', ']'},
		{'<', '>'},
	}

	openingCharStack := []rune{}

	for _, c := range line {
		if d.isOpeningChar(acceptedPairs, c) {
			openingCharStack = append(openingCharStack, c)
			continue
		}

		if d.isClosingChar(acceptedPairs, c) {
			// Pop
			n := len(openingCharStack) - 1
			lastOpening := openingCharStack[n]
			openingCharStack = openingCharStack[:n]
			expected, _ := d.getClosingChar(acceptedPairs, lastOpening)
			if expected != c {
				fmt.Printf("Expected %v but found %v instead", string(expected), string(c))
			}

		}
	}
	// Pop
	// n := len(stack) - 1
	// basinPoint := stack[n]
	// stack = stack[:n]
}

// make all of these methods to some struct with slice of pairs?
func (DayTen) isOpeningChar(acceptedPairs []tagPair, char rune) bool {
	for _, pair := range acceptedPairs {
		if pair.opening == char {
			return true
		}
	}
	return false
}

func (DayTen) isClosingChar(acceptedPairs []tagPair, char rune) bool {
	for _, pair := range acceptedPairs {
		if pair.closing == char {
			return true
		}
	}
	return false
}

func (DayTen) getClosingChar(acceptedPairs []tagPair, char rune) (rune, bool) {
	for _, pair := range acceptedPairs {
		if pair.opening == char {
			return pair.closing, true
		}
	}
	return 0, false
}

func (d DayTen) PartTwo() string {
	return "Not implemented"
}

func (d *DayTen) init() {
}
