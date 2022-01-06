package puzzles

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DayEight struct {
	day
	entries              []segmentDisplayEntry
	segmentLengthToDigit map[int]int
	segmentBinToInt      map[int]int // according to PGFEDCBA format
}

type segmentDisplayEntry struct {
	uniqueSignalPatterns [10]segmentPattern // 10 unique seven-segment patterns
	outputValue          [4]segmentPattern  // 4 seven-segment digits
}

type segmentPattern struct {
	pattern string
}

func (DayEight) GetPuzzleName() string {
	return "Day 8: Seven Segment Search"
}

func (d DayEight) PartOne() string {
	// How many times do digits 1, 4, 7 or 8 appear?
	// seven-segment 1 has signal length = 2, 4 = 4, 7 = 3 & 8 = 7
	var sum1478 int
	for _, entry := range d.entries {
		for _, segmentPattern := range entry.outputValue {
			if _, exists := d.segmentLengthToDigit[segmentPattern.length()]; exists {
				sum1478++
			}
		}
	}

	return fmt.Sprintf("1,4,7 or 8 appears %d times", sum1478)
}

func (d DayEight) PartTwo() string {
	sum := 0
	for _, displayEntry := range d.entries {
		cipher := d.getCipherFromSignalPatterns(displayEntry.uniqueSignalPatterns[:])
		var number string
		for _, outputValue := range displayEntry.outputValue {
			digit := d.decodeNumber(outputValue, cipher)
			number += fmt.Sprintf("%d", digit)
		}

		sum += parseInt(number)
	}

	return fmt.Sprintf("The sum of all output values is %d", sum)
}

// cipher is in PGFEDCBA format, each char at index corresponds to segment position
//0 dot (unused)
//1 middle G
//2 top left F
//3 bottom left E
//4 bottom D
//5 bottom right C
//6 top right B
//7 top A
func (d DayEight) getCipherFromSignalPatterns(signalPatterns []segmentPattern) [8]rune {
	// get pattern for 1, 4, 7 & 8
	initialMap := d.getInitialMappings(signalPatterns)
	one := initialMap[1]
	four := initialMap[4]
	seven := initialMap[7]
	eight := initialMap[8]

	cipher := [8]rune{}

	// build cypher in a bruteforcey manner
	for _, char := range seven.pattern {
		if !one.contains(char) {
			cipher[7] = char
		}
	}
	for _, segment := range signalPatterns {
		if segment.length() == 6 { // 9, 6, 0
			for _, char := range eight.pattern {
				if !segment.contains(char) && one.contains(char) {
					cipher[6] = char
					cipher[5] = rune(strings.Replace(one.pattern, string(char), "", 1)[0])
				} else if !segment.contains(char) && four.contains(char) {
					cipher[1] = char
				} else if !segment.contains(char) {
					cipher[3] = char
				}
			}
		}
	}
	for _, char := range four.pattern {
		if char != cipher[6] && char != cipher[5] && char != cipher[1] {
			cipher[2] = char
		}
	}
	for _, char := range eight.pattern {
		contains := false
		for _, a := range cipher {
			if char == a {
				contains = true
			}
		}
		if !contains {
			cipher[4] = char
			break
		}
	}

	return cipher
}

func (d DayEight) decodeNumber(signalPattern segmentPattern, cipher [8]rune) int {
	var binString = ""
	for _, c := range cipher {
		if signalPattern.contains(c) {
			binString += "1"
		} else {
			binString += "0"
		}
	}
	// d.printNumber(binString, cipher)
	if i, err := strconv.ParseInt(binString, 2, 16); err != nil {
		panic(err)
	} else {
		digit, exists := d.segmentBinToInt[int(i)]
		if !exists {
			panic("unmapped segment")
		}
		return digit
	}
}

// func (DayEight) printNumber(binString string, cipher [8]rune) {
// 	p := printer{2, 4, 2}

// 	display := p.getDisplay([]string{strings.Replace(binString, "0", " ", -1)})
// 	for _, row := range display {
// 		fmt.Println(row)
// 	}
// }

func (d DayEight) getInitialMappings(signalPatterns []segmentPattern) map[int]segmentPattern {
	patternToDigit := make(map[int]segmentPattern)
	for _, pattern := range signalPatterns {
		if !pattern.isAmbiguous() {
			digit, _ := pattern.getCorrespondingDigit(d.segmentLengthToDigit)
			patternToDigit[digit] = pattern
		}
	}
	return patternToDigit
}

func (d *DayEight) init() {
	for _, line := range d.inputLines {
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

	d.initMaps()
}

func (d *DayEight) initMaps() {
	d.segmentLengthToDigit = make(map[int]int)
	d.segmentLengthToDigit[2] = 1
	d.segmentLengthToDigit[7] = 8
	d.segmentLengthToDigit[3] = 7
	d.segmentLengthToDigit[4] = 4

	// int values of binary representation of segments
	d.segmentBinToInt = map[int]int{
		63:  0, // 00111111
		6:   1, // 00000110
		91:  2, // 01011011
		79:  3, // 01001111
		102: 4, // 01100110
		109: 5, // 01101101
		125: 6, // 01111101
		7:   7, // 00000111
		127: 8, // 01111111
		111: 9, // 01101111
	}
}

func (s segmentPattern) length() int {
	return len(s.pattern)
}

func (s segmentPattern) contains(char rune) bool {
	return strings.ContainsRune(s.pattern, char)
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
