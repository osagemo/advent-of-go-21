package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

type DayTwo struct {
	day
	acceptedDirections map[string]bool
}

type subMarineState struct {
	xPos int
	yPos int
	aim  int
}

func (DayTwo) GetPuzzleName() string {
	return "Day 2: Dive!"
}

func (d *DayTwo) init() {
	d.acceptedDirections = map[string]bool{"forward": true, "up": true, "down": true}
}

func (d DayTwo) PartOne() string {
	sub := subMarineState{0, 0, 0}

	for _, command := range d.input.Lines {
		sub = d.parseCommand(sub, command, 1)
	}

	solution := fmt.Sprintf("Horizontal position: %d, Depth: %d. Solution: %d", sub.xPos, sub.yPos, sub.xPos*sub.yPos)
	return solution
}

func (d DayTwo) PartTwo() string {
	sub := subMarineState{0, 0, 0}

	for _, command := range d.input.Lines {
		sub = d.parseCommand(sub, command, 2)
	}

	solution := fmt.Sprintf("Horizontal position: %d, Depth: %d. Solution: %d", sub.xPos, sub.yPos, sub.xPos*sub.yPos)
	return solution
}

func (d DayTwo) parseCommand(initialState subMarineState, command string, part int) (newState subMarineState) {
	parts := strings.Split(command, " ")

	if len(parts) != 2 {
		panic(fmt.Sprintf("bad input: %s", command))
	}

	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("bad input: %s", command))
	}

	direction := parts[0]
	if !d.acceptedDirections[direction] {
		panic(fmt.Sprintf("unrecognized direction: %s", command))
	}

	newState = initialState.move(direction, amount, part)
	return newState
}

// Hard-coded strings are used instead of enum but validated against acceptedDirections map
func (state subMarineState) move(direction string, amount int, part int) subMarineState {
	switch direction {
	case "forward":
		state.changeX(amount, part)
	case "up":
		state.changeY(-amount, part)
	case "down":
		state.changeY(amount, part)
	}

	return state
}

func (state *subMarineState) changeX(amount int, part int) {
	if part == 1 {
		state.xPos += amount
	} else if part == 2 {
		state.xPos += amount
		state.yPos += (state.aim * amount)
	}
}

func (state *subMarineState) changeY(amount int, part int) {
	if part == 1 {
		state.yPos += amount
	} else if part == 2 {
		state.aim += amount
	}
}
