package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

type DayThree struct {
	day
	zeros [12]int
	ones  [12]int
}

func (DayThree) GetPuzzleName() string {
	return "Day 3: Binary Diagnostic"
}

func (d *DayThree) PartOne() string {
	var gammaBin = ""
	var epsilonBin = ""
	for i := 0; i < 12; i++ {
		if d.ones[i] > d.zeros[i] {
			gammaBin += "1"
			epsilonBin += "0"
		} else {
			gammaBin += "0"
			epsilonBin += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBin, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonBin, 2, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Gamma: %v (%v), Epsilon: %v (%v), Solution: %v", gamma, gammaBin, epsilon, epsilonBin, gamma*epsilon)
}

func (d *DayThree) PartTwo() string {
	oxygenRatingBin := d.getOxygenRating()
	co2ScrubberRatingBin := d.getCO2ScrubberRating()

	oxygenRating, err := strconv.ParseInt(oxygenRatingBin, 2, 64)
	if err != nil {
		panic(err)
	}

	co2ScrubberRating, err := strconv.ParseInt(co2ScrubberRatingBin, 2, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("oxygen: %d, co2: %d, solution: %d", oxygenRating, co2ScrubberRating, oxygenRating*co2ScrubberRating)
}

func (d *DayThree) getOxygenRating() string {
	return d.findValue(true)
}

func (d *DayThree) getCO2ScrubberRating() string {
	return d.findValue(false)
}

func (d *DayThree) findValue(highest bool) string {
	var consideredLines = d.inputLines
	var index = 0
	var wantedValue string
	// copy frequency counts
	zeros := d.zeros
	ones := d.ones

	for len(consideredLines) > 1 {
		wantedValue = d.getWantedValue(ones, zeros, highest, index)
		consideredLines, zeros, ones = reduceDiagnosticLines(consideredLines, wantedValue, index)
		index++
	}

	return strings.Join(consideredLines, "")
}

func (DayThree) getWantedValue(ones [12]int, zeros [12]int, highest bool, index int) string {
	var initialValue string
	if ones[index] >= zeros[index] {
		if highest {
			initialValue = "1"
		} else {
			initialValue = "0"
		}
	} else {
		if highest {
			initialValue = "0"
		} else {
			initialValue = "1"
		}
	}
	return initialValue
}

func reduceDiagnosticLines(lines []string, initialVal string, charIndex int) ([]string, [12]int, [12]int) {
	var zeros = [12]int{}
	var ones = [12]int{}
	consideredLines := []string{}

	for _, line := range lines {
		if len(line) <= charIndex {
			panic("charIndex outside of bounds")
		}
		firstChar := line[charIndex : charIndex+1]

		if firstChar == initialVal {
			consideredLines = append(consideredLines, line)
			for i, char := range line {
				if char == 48 {
					zeros[i]++
				} else if char == 49 {
					ones[i]++
				}
			}
		}
	}

	return consideredLines, zeros, ones
}

// inefficient (iterating over input multiple times) but convenient
func (d *DayThree) init() {
	for _, line := range d.inputLines {
		for i, char := range line {
			if char == 48 {
				d.zeros[i]++
			} else if char == 49 {
				d.ones[i]++
			}
		}
	}
}
