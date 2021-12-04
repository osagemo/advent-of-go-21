package puzzles

import (
	"fmt"
	"math"
	"strconv"
)

type DayOne struct {
	base         day
	measurements []int
}

// inefficient but convenient
func (d *DayOne) init() {
	for _, line := range d.base.input.Lines {
		measurement, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		d.measurements = append(d.measurements, measurement)
	}
}

func (d *DayOne) GetPuzzleName() string {
	return "Day 1: Sonar Sweep"
}

func (d *DayOne) PartOne() string {
	numOfIncreases := d.getNumOfIncreasingMeasurements()

	solution := fmt.Sprintf("The number of measurements larger than the previous measurement is: %d", numOfIncreases)
	return solution
}

func (d *DayOne) PartTwo() string {
	numOfIncreasingWindows := d.getNumOfIncreasingMeasurementWindows(3)

	solution := fmt.Sprintf("The number of sums larger than the previous sums is: %d", numOfIncreasingWindows)
	return solution
}

func (d *DayOne) getNumOfIncreasingMeasurements() int {
	increases := 0
	prev := math.MaxInt32

	for _, m := range d.measurements {
		if m > prev {
			increases++
		}
		prev = m
	}
	return increases
}

func (d *DayOne) getNumOfIncreasingMeasurementWindows(windowSize int) int {
	windowIndex := windowSize - 1
	increases := 0
	currentSum := 0
	prevSum := math.MaxInt32

	for i, measurement := range d.measurements {
		currentSum += measurement

		if i >= windowIndex {
			// Compare sum to previous
			if currentSum > prevSum {
				increases++
			}
			prevSum = currentSum

			// Remove first element in window from sum
			currentSum -= d.measurements[i-windowIndex]
		}
	}
	return increases
}
