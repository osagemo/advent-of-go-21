package puzzles

import (
	"errors"
	"fmt"
	"strings"
)

type DayEight struct {
	day
	entries              []segmentDisplayEntry
	segmentLengthToDigit map[int]int
}

type segmentDisplayEntry struct {
	uniqueSignalPatterns [10]segmentPattern // 10 unique seven-segment patterns
	outputValue          [4]segmentPattern  // 4 seven-segment digits
}

type segmentPattern struct {
	pattern string
}

func (s segmentPattern) length() int {
	return len(s.pattern)
}

func (s segmentPattern) isAmbiguous() bool {
	unqiueSegmentLengths := []int{2, 4, 3, 7}
	return !contains(unqiueSegmentLengths, s.length())
}

func (s segmentPattern) getCorrespondingDigit(segmentLengthMap map[int]int) (int, error) {
	if s.isAmbiguous() {
		return -1, errors.New("pattern is ambiguous")
	}

	digit, exists := segmentLengthMap[s.length()]
	if exists {
		return digit, nil
	}

	return -1, errors.New("unmapped digit")
}

func (DayEight) GetPuzzleName() string {
	return "Day 8: Seven Segment Search"
}

func (d DayEight) PartOne() string {
	// How many times do digits 1, 4, 7 or 8 appear?
	// seven-segment 1 has length = 2, 4 = 4, 7 = 3 & 8 = 7
	// segmentLengthFrequencies := make(map[int]int)
	var sum1478 int
	for _, entry := range d.entries {
		for _, segmentPattern := range entry.outputValue {
			// segmentLengthFrequencies[segmentPattern.length()]++
			if _, exists := d.segmentLengthToDigit[segmentPattern.length()]; exists {
				sum1478++
			}
		}
	}

	return fmt.Sprintf("1,4,7 or 8 appears %d times", sum1478)
}

func (d *DayEight) PartTwo() string {
	p := printer{2, 6, 4, 2}
	testDigits := []string{"111111 ", "11     ", "1 11 11", "1  1111", "11  11 ", "11 11 1", "11111 1", "111    ", "1111111", "11 1111", "111 111", "11111  ", "111  1 ", "1 1111 ", "1111  1", "111   1"}

	p.test(testDigits)
	return "test"
}

func (d DayEight) mapSignalPatternsToDigits(signalPatterns []segmentPattern) map[segmentPattern]int {
	// get pattern for 1, 4, 7 & 8
	patternToDigit := make(map[segmentPattern]int)
	for _, pattern := range signalPatterns {
		if !pattern.isAmbiguous() {
			digit, _ := pattern.getCorrespondingDigit(d.segmentLengthToDigit)
			patternToDigit[pattern] = digit
		}
	}
	return patternToDigit
}

func (DayEight) signalPatternToDigit(signal [8]string) {

}

type printer struct {
	digitSeparatorLength int
	digitWidth           int
	horizontalWidth      int
	verticalHeight       int
}

func (p printer) middle(display [7]string, digit string, row int) [7]string {
	middle := 7
	// bs := []byte(display[row])
	if row == 3 {
		middle = 1
	} else if row == 7 {
		middle = 4
	}

	// bs = append(bs, ' ')
	display[row] += " "
	for i := 0; i < p.horizontalWidth; i++ {
		// bs = append(bs, digit[middle])
		display[row] += string(digit[middle])
	}
	// bs = append(bs, ' ')
	display[row] += " "

	// display[row] = string(bs)
	return display
}

func (p printer) leftRight(display [7]string, digit string, row int) [7]string {
	left := 2
	right := 6

	if row > 3 {
		left = 3
		right = 5
	}
	bs := []byte{}
	bs[row] += digit[left]
	for i := 0; i < p.horizontalWidth; i++ {
		bs[row] += ' '
	}
	bs[row] += digit[right]

	display[row] = string(bs)
	return display
}

func (p printer) appendDigit(display [7]string, digit string) [7]string {
	display = p.middle(display, digit, 0)
	display = p.leftRight(display, digit, 1)
	display = p.leftRight(display, digit, 2)
	display = p.middle(display, digit, 3)
	display = p.leftRight(display, digit, 4)
	display = p.leftRight(display, digit, 5)
	display = p.middle(display, digit, 6)
	return display
}

// (p)GFEDCBA
func (p printer) test(digits []string) {
	totalLength := p.digitWidth * len(digits)

	var display = [7][totalLength]byte

	for _, digit := range digits {
		p.appendDigit(display, digit)
	}

	//0 dot

	//1 middle G

	//2 top left F

	//3 bottom left E

	//4 bottom D

	//5 bottom right C

	//6 top right B

	//7 top A
	for _, row := range display {
		fmt.Println(row)
	}
}

func (d *DayEight) init() {
	for _, line := range d.input.Lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			panic("bad input")
		}

		var signalPatterns [10]segmentPattern
		var outputValue [4]segmentPattern
		for i, pattern := range strings.Fields(parts[0]) {
			signalPatterns[i] = segmentPattern{pattern}
		}
		for i, pattern := range strings.Fields(parts[1]) {
			outputValue[i] = segmentPattern{pattern}
		}
		d.entries = append(d.entries, segmentDisplayEntry{signalPatterns, outputValue})
	}

	d.initSegmentLengthMap()
}

func (d *DayEight) initSegmentLengthMap() {
	d.segmentLengthToDigit = make(map[int]int)
	d.segmentLengthToDigit[2] = 1
	d.segmentLengthToDigit[7] = 8
	d.segmentLengthToDigit[3] = 7
	d.segmentLengthToDigit[4] = 4
}
