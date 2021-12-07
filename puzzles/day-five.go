package puzzles

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type DayFive struct {
	day
	lines []line
}

func (DayFive) GetPuzzleName() string {
	return "Day 5: Hydrothermal Venture"
}

func (d *DayFive) PartOne() string {
	coordinateToVentLines := d.mapLines(d.lines, false)

	sum := 0
	for _, v := range coordinateToVentLines {
		if v > 1 {
			sum++
		}
	}

	return fmt.Sprintf("Sum of coordinates with more than one line covering it: %d", sum)

}

func (d *DayFive) PartTwo() string {
	coordinateToVentLines := d.mapLines(d.lines, true)

	sum := 0
	for _, v := range coordinateToVentLines {
		if v > 1 {
			sum++
		}
	}

	return fmt.Sprintf("Sum of coordinates with more than one line covering it: %d", sum)
}

func (d *DayFive) init() {
	re := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	for _, string := range d.input.Lines {
		if re.MatchString(string) {
			matches := re.FindStringSubmatch(string)
			from := coordinate{parseInt(matches[1]), parseInt(matches[2])}
			to := coordinate{parseInt(matches[3]), parseInt(matches[4])}
			newLine := line{from, to}
			d.lines = append(d.lines, newLine)
		}
	}
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
	return math.Abs(float64(l.from.x-l.to.x)) == math.Abs(float64(l.from.y-l.to.y))
}

func (l line) isEven() bool {
	return l.isHorizontal() || l.isVertical()
}

func (l line) train() []coordinate {
	train := []coordinate{}
	if l.isVertical() {
		var min, max int
		if l.from.x < l.to.x {
			min = l.from.x
			max = l.to.x
		} else {
			min = l.to.x
			max = l.from.x
		}
		for i := min; i <= max; i++ {
			train = append(train, coordinate{i, l.from.y})
		}
		return train
	}

	if l.isHorizontal() {
		var min, max int
		if l.from.y < l.to.y {
			min = l.from.y
			max = l.to.y
		} else {
			min = l.to.y
			max = l.from.y
		}
		for i := min; i <= max; i++ {
			train = append(train, coordinate{l.from.x, i})
		}
		return train
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
			train = append(train, coordinate{x, y})
			y += yDif
		}
		train = append(train, coordinate{l.to.x, l.to.y})
	}

	return train
}

func (DayFive) mapLines(lines []line, includeDiagonal bool) map[coordinate]int {
	lineMap := make(map[coordinate]int)
	for _, l := range lines {
		if l.isEven() {
			for _, coord := range l.train() {
				lineMap[coord]++
			}
		} else if includeDiagonal && l.isDiagonal() {
			for _, coord := range l.train() {
				lineMap[coord]++
			}
		}
	}
	return lineMap
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
