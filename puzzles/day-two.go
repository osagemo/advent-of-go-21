package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

// I wanted to use Enums but had a hard time wrapping my head around golangs version of them..
// move out of global scope? to init function?
var acceptedDirections = map[string]bool{"forward": true, "up": true, "down": true}

type DayTwo struct {
	base day
	xPos int
	yPos int
	aim  int
}

func (d *DayTwo) GetPuzzleName() string {
	return "Day 2: Dive!"
}

func (d *DayTwo) PartOne() string {
	for _, command := range d.base.input.Lines {
		d.parseCommand(command, 1)
	}

	solution := fmt.Sprintf("Horizontal position: %d, Depth: %d. Solution: %d", d.xPos, d.yPos, d.xPos*d.yPos)
	return solution
}

func (d *DayTwo) PartTwo() string {
	d.init() // reset

	for _, command := range d.base.input.Lines {
		d.parseCommand(command, 2)
	}

	solution := fmt.Sprintf("Horizontal position: %d, Depth: %d. Solution: %d", d.xPos, d.yPos, d.xPos*d.yPos)
	return solution
}

func (d *DayTwo) init() {
	d.xPos = 0
	d.yPos = 0
	d.aim = 0
}

func (d *DayTwo) parseCommand(command string, part int) {
	parts := strings.Split(command, " ")

	if len(parts) != 2 {
		panic(fmt.Sprintf("bad input: %s", command))
	}

	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("bad input: %s", command))
	}

	direction := parts[0]
	if !acceptedDirections[direction] {
		panic(fmt.Sprintf("unrecognized direction: %s", command))
	}

	d.move(direction, amount, part)
}

// Hard-coded strings are used instead of enum but validated against acceptedDirections map
func (d *DayTwo) move(direction string, amount int, part int) {
	switch direction {
	case "forward":
		d.changeX(amount, part)
	case "up":
		d.changeY(-amount, part)
	case "down":
		d.changeY(amount, part)
	}
}

func (d *DayTwo) changeX(amount int, part int) {
	if part == 1 {
		d.xPos += amount
	} else if part == 2 {
		d.xPos += amount
		d.yPos += (d.aim * amount)
	}
}

func (d *DayTwo) changeY(amount int, part int) {
	if part == 1 {
		d.yPos += amount
	} else if part == 2 {
		d.aim += amount
	}
}
