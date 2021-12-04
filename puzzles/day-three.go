package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

type DayThree struct {
	base        day
	zeros       [12]int
	ones        [12]int
	reportLines [][12]int
}

func (d *DayThree) GetPuzzleName() string {
	return "Day 3: Binary Diagnostic"
}

func (d *DayThree) PartOne() string {
	var gammaBin = make([]string, 12)
	var epsilonBin = make([]string, 12)
	for i := 0; i < 12; i++ {
		if d.ones[i] > d.zeros[i] {
			gammaBin[i] = "1"
			epsilonBin[i] = "0"
		} else {
			gammaBin[i] = "0"
			epsilonBin[i] = "1"
		}
	}

	gamma, err := strconv.ParseInt(strings.Join(gammaBin, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(strings.Join(epsilonBin, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Gamma: %v (%v), Epsilon: %v (%v), Solution: %v", gamma, gammaBin, epsilon, epsilonBin, gamma*epsilon)
}

func (d *DayThree) PartTwo() string {
	return "Not implemented yet"
}

// inefficient but convenient
func (d *DayThree) init() {
	for _, line := range d.base.input.Lines {
		for i, char := range line {
			if char == 48 {
				d.zeros[i]++
			} else if char == 49 {
				d.ones[i]++
			}
		}
	}
}
