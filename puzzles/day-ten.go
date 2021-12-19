package puzzles

import (
	"errors"
	"fmt"
	"sort"
)

type DayTen struct {
	day
	format formatRules
}

type tagPair struct {
	opening rune
	closing rune
}

type formatRules struct {
	acceptedPairs []tagPair
}

func (DayTen) GetPuzzleName() string {
	return "Day 10: Syntax Scoring"
}

func (d DayTen) PartOne() string {
	score := d.getCorruptedScore(d.input.Lines)
	return fmt.Sprintf("Total syntax error score: %d", score)
}

func (d DayTen) getCorruptedScore(lines []string) int {
	score := 0
	scoreTable := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, line := range lines {
		isCorrupted, _, found := d.lineIsCorrupted(line)
		if isCorrupted {
			score += scoreTable[found]
		}
	}
	return score
}

func (d DayTen) PartTwo() string {
	scores := d.getAutoCompleteScores(d.input.Lines)
	sort.Ints(scores)
	middleScore := scores[len(scores)/2]
	return fmt.Sprintf("The middle score is %d\n", middleScore)
}

func (d DayTen) getAutoCompleteScores(lines []string) []int {
	scores := []int{}

	for _, line := range lines {
		closingString, err := d.getClosingChars(line)
		if err == nil {
			scores = append(scores, d.getCompletionStringScore(closingString))
		}
	}
	return scores
}

func (d DayTen) getCompletionStringScore(line string) int {
	score := 0
	scoreTable := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	for _, c := range line {
		score *= 5
		score += scoreTable[c]
	}

	return score
}

func (d DayTen) lineIsCorrupted(line string) (isCorrupted bool, expected rune, found rune) {
	found = 0
	expected = 0
	isCorrupted = true
	openingCharStack := []rune{}

	for _, c := range line {
		if d.format.isOpeningChar(c) {
			openingCharStack = append(openingCharStack, c)
			continue
		}

		if d.format.isClosingChar(c) {
			// Pop
			n := len(openingCharStack) - 1
			lastOpening := openingCharStack[n]
			openingCharStack = openingCharStack[:n]
			want, _ := d.format.getClosingChar(lastOpening)
			if want != c {
				found = c
				expected = want
				return
			}
		}
	}
	isCorrupted = false
	return
}

func (d DayTen) getClosingChars(line string) (string, error) {
	completionString := ""
	openingCharStack := []rune{}

	for _, c := range line {
		if d.format.isOpeningChar(c) {
			openingCharStack = append(openingCharStack, c)
			continue
		}

		if d.format.isClosingChar(c) {
			// Pop
			n := len(openingCharStack) - 1
			lastOpening := openingCharStack[n]
			openingCharStack = openingCharStack[:n]
			want, _ := d.format.getClosingChar(lastOpening)
			if want != c {
				return completionString, errors.New("line is corrupted")
			}
		}
	}

	for i := len(openingCharStack) - 1; i >= 0; i-- {
		closing, _ := d.format.getClosingChar(openingCharStack[i])
		completionString += string(closing)
	}
	return completionString, nil
}

func (f formatRules) isOpeningChar(char rune) bool {
	for _, pair := range f.acceptedPairs {
		if pair.opening == char {
			return true
		}
	}
	return false
}

func (f formatRules) isClosingChar(char rune) bool {
	for _, pair := range f.acceptedPairs {
		if pair.closing == char {
			return true
		}
	}
	return false
}

func (f formatRules) getClosingChar(char rune) (rune, bool) {
	for _, pair := range f.acceptedPairs {
		if pair.opening == char {
			return pair.closing, true
		}
	}
	return 0, false
}

func (d *DayTen) init() {
	d.format = formatRules{[]tagPair{
		{'{', '}'},
		{'(', ')'},
		{'[', ']'},
		{'<', '>'},
	}}
}
