package puzzles

import (
	"fmt"
	"regexp"
)

type DayFive struct {
	day
	ventLines []line
}

func (DayFive) GetPuzzleName() string {
	return "Day 5: Hydrothermal Venture"
}

func (d *DayFive) PartOne() string {
	coordsCovered := d.getCoordinateFrequencyMap(d.ventLines, false)

	sum := 0
	for _, v := range coordsCovered {
		if v > 1 {
			sum++
		}
	}

	return fmt.Sprintf("Sum of coordinates with more than one vent covering it: %d", sum)
}

func (d *DayFive) PartTwo() string {
	coordsCovered := d.getCoordinateFrequencyMap(d.ventLines, true)

	sum := 0
	for _, v := range coordsCovered {
		if v > 1 {
			sum++
		}
	}

	return fmt.Sprintf("Sum of coordinates with more than one vent covering it: %d", sum)
}

func (d *DayFive) init() {

	for _, string := range d.inputLines {
		line := d.parseLine(string)
		d.ventLines = append(d.ventLines, line)
	}
}

func (DayFive) parseLine(s string) line {
	newLine := line{}
	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	if re.MatchString(s) {
		matches := re.FindStringSubmatch(s)
		from := coordinate{parseInt(matches[1]), parseInt(matches[2])}
		to := coordinate{parseInt(matches[3]), parseInt(matches[4])}
		newLine = line{from, to}
	} else {
		panic("invalid input")
	}

	return newLine
}

type line struct {
	from coordinate
	to   coordinate
}

func (l line) isHorizontal() bool {
	return l.from.x == l.to.x
}

func (l line) isVertical() bool {
	return l.from.y == l.to.y
}

func (l line) isDiagonal() bool {
	return absDiffInt(l.from.x, l.to.x) == absDiffInt(l.from.y, l.to.y)
}

func (l line) isEven() bool {
	return l.isHorizontal() || l.isVertical()
}

func (l line) getIncludedCoordinates() []coordinate {
	coordinateTrain := []coordinate{}
	if l.isVertical() {
		var min = l.to.x
		var max = l.from.x
		if l.from.x < l.to.x {
			min = l.from.x
			max = l.to.x
		}
		for i := min; i <= max; i++ {
			coordinateTrain = append(coordinateTrain, coordinate{i, l.from.y})
		}
		return coordinateTrain
	}

	if l.isHorizontal() {
		var min = l.to.y
		var max = l.from.y
		if l.from.y < l.to.y {
			min = l.from.y
			max = l.to.y
		}
		for i := min; i <= max; i++ {
			coordinateTrain = append(coordinateTrain, coordinate{l.from.x, i})
		}
		return coordinateTrain
	}

	if l.isDiagonal() {
		xDif, yDif := 1, 1 // x,y are increasing
		if l.from.x > l.to.x {
			xDif = -1 // x is decreasing
		}
		if l.from.y > l.to.y {
			yDif = -1 // y is decreasing
		}

		y := l.from.y
		for x := l.from.x; x != l.to.x; x += xDif {
			coordinateTrain = append(coordinateTrain, coordinate{x, y})
			y += yDif
		}
		coordinateTrain = append(coordinateTrain, coordinate{l.to.x, l.to.y})
	}

	return coordinateTrain
}

func (DayFive) getCoordinateFrequencyMap(lines []line, includeDiagonal bool) map[coordinate]int {
	coordFrequencies := make(map[coordinate]int)
	for _, l := range lines {
		if l.isEven() {
			for _, coord := range l.getIncludedCoordinates() {
				coordFrequencies[coord]++
			}
		} else if includeDiagonal && l.isDiagonal() {
			for _, coord := range l.getIncludedCoordinates() {
				coordFrequencies[coord]++
			}
		}
	}
	return coordFrequencies
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
